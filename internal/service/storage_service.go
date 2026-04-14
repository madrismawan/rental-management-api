package service

import (
	"context"
	"fmt"
	"mime"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"rental-management-api/config"
)

type StorageService interface {
	Upload(ctx context.Context, file *multipart.FileHeader, prefix string) (string, error)
	ResolveURL(objectRef string) (string, error)
}

type storageService struct {
	client *s3.S3
	bucket string
}

func NewStorageService(cfg config.StorageConfig) StorageService {
	endpoint, secure := normalizeEndpoint(cfg.Endpoint)
	awsCfg := aws.Config{
		Region:           aws.String(cfg.Region),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	}

	if endpoint != "" {
		awsCfg.Endpoint = aws.String(endpoint)
		awsCfg.DisableSSL = aws.Bool(!secure)
	}

	sess, err := session.NewSession(&awsCfg)
	if err != nil {
		return &storageService{bucket: cfg.Bucket}
	}

	return &storageService{
		client: s3.New(sess),
		bucket: cfg.Bucket,
	}
}

func (s *storageService) Upload(ctx context.Context, file *multipart.FileHeader, prefix string) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file is required")
	}
	if s.client == nil || s.bucket == "" {
		return "", fmt.Errorf("storage configuration is incomplete")
	}

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("open upload file: %w", err)
	}
	defer src.Close()

	objectName := buildObjectName(prefix, file.Filename)
	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = mime.TypeByExtension(filepath.Ext(file.Filename))
	}
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err = s.client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(objectName),
		Body:        src,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("put object: %w", err)
	}

	return objectName, nil
}

func (s *storageService) ResolveURL(objectRef string) (string, error) {
	if objectRef == "" {
		return "", nil
	}
	if s.client == nil || s.bucket == "" {
		return objectRef, nil
	}

	objectName := extractObjectName(objectRef, s.bucket)
	if objectName == "" {
		return objectRef, nil
	}

	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(objectName),
	})
	presignedURL, err := req.Presign(24 * time.Hour)
	if err != nil {
		return "", fmt.Errorf("presign object: %w", err)
	}

	return presignedURL, nil
}

func buildObjectName(prefix string, filename string) string {
	cleanPrefix := strings.Trim(prefix, "/")
	cleanFilename := strings.ReplaceAll(strings.TrimSpace(filename), " ", "-")
	if cleanFilename == "" {
		cleanFilename = "file"
	}

	if cleanPrefix == "" {
		return fmt.Sprintf("%d-%s", time.Now().UnixNano(), cleanFilename)
	}

	return fmt.Sprintf("%s/%d-%s", cleanPrefix, time.Now().UnixNano(), cleanFilename)
}

func normalizeEndpoint(raw string) (endpoint string, secure bool) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", true
	}

	if strings.HasPrefix(trimmed, "http://") || strings.HasPrefix(trimmed, "https://") {
		u, err := url.Parse(trimmed)
		if err == nil && u.Host != "" {
			return strings.TrimRight(u.Scheme+"://"+u.Host, "/"), u.Scheme == "https"
		}
	}

	cleaned := strings.Trim(trimmed, "/")
	return "https://" + cleaned, true
}

func extractObjectName(objectRef string, bucket string) string {
	ref := strings.TrimSpace(objectRef)
	if ref == "" {
		return ""
	}

	if strings.HasPrefix(ref, "s3://") {
		withoutScheme := strings.TrimPrefix(ref, "s3://")
		parts := strings.SplitN(withoutScheme, "/", 2)
		if len(parts) == 2 {
			if parts[0] == bucket {
				return strings.Trim(parts[1], "/")
			}
			return ""
		}
		return ""
	}

	if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") {
		u, err := url.Parse(ref)
		if err != nil {
			return ""
		}
		path := strings.Trim(u.Path, "/")
		if path == "" {
			return ""
		}
		if strings.HasPrefix(path, bucket+"/") {
			return strings.TrimPrefix(path, bucket+"/")
		}
		return ""
	}

	return strings.Trim(ref, "/")
}
