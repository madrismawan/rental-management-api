package service

import (
	"errors"
	"fmt"
	"rental-management-api/internal/constant"
	"rental-management-api/internal/dto"
	"rental-management-api/internal/entity"
	"rental-management-api/internal/mapper"
	"rental-management-api/pkg/errs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(c *gin.Context, input dto.RegisterRequest) (entity.User, error)
	Login(c *gin.Context, email string, password string) (entity.User, dto.AuthTokenResource, error)
	VerifyToken(c *gin.Context, token string) (entity.User, error)
}

type authService struct {
	userService        UserService
	accessTokenSecret  []byte
	refreshTokenSecret []byte
	ttl                time.Duration
}

type AuthClaims struct {
	jwt.RegisteredClaims
	User dto.UserResource
}

func NewAuthService(userService UserService, accessTokenSecret string, refreshTokenSecret string, ttl time.Duration) AuthService {
	return &authService{
		userService:        userService,
		accessTokenSecret:  []byte(accessTokenSecret),
		refreshTokenSecret: []byte(refreshTokenSecret),
		ttl:                ttl,
	}
}

func (s *authService) Register(c *gin.Context, input dto.RegisterRequest) (entity.User, error) {
	u, err := s.userService.Create(c, CreateUserInput{
		Name:     input.Name,
		Email:    input.Email,
		Role:     constant.UserRoleCustomer,
		Password: input.Password,
	})
	if err != nil {
		return entity.User{}, err
	}
	u.Password = ""
	return *u, nil
}

func (s *authService) Login(c *gin.Context, email string, password string) (entity.User, dto.AuthTokenResource, error) {
	var u entity.User
	if email == "" || password == "" {
		return u, dto.AuthTokenResource{}, errs.ErrUserNotFound
	}

	u, err := s.userService.GetByColumn(c, "email", email)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return u, dto.AuthTokenResource{}, errs.ErrInvalidCredentials
		}
		return u, dto.AuthTokenResource{}, err
	}

	if err := comparePassword(u.Password, password); err != nil {
		return u, dto.AuthTokenResource{}, errs.ErrInvalidCredentials
	}
	token, err := s.generateAccessToken(u)
	if err != nil {
		return u, dto.AuthTokenResource{}, fmt.Errorf("generate access token: %w", err)
	}
	refreshToken, err := s.generateRefreshToken(u)
	if err != nil {
		return u, dto.AuthTokenResource{}, fmt.Errorf("generate refresh token: %w", err)
	}

	return u, dto.AuthTokenResource{
		AccessToken:  token,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}, nil
}

func (s *authService) VerifyToken(c *gin.Context, token string) (entity.User, error) {
	if token == "" {
		return entity.User{}, errs.ErrInvalidCredentials
	}

	claims := &AuthClaims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.accessTokenSecret, nil
	})
	if err != nil {
		return entity.User{}, errs.ErrInvalidCredentials
	}

	if !parsed.Valid {
		return entity.User{}, errs.ErrInvalidCredentials
	}

	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return entity.User{}, errs.ErrInvalidCredentials
	}
	user, err := s.userService.GetByColumn(c, "id", userID)
	if err != nil {
		return entity.User{}, errs.ErrInvalidCredentials
	}

	return user, nil
}

func comparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (s *authService) generateAccessToken(user entity.User) (string, error) {
	claims := AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: mapper.ToUserResource(user),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.accessTokenSecret)
}

func (s *authService) generateRefreshToken(user entity.User) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.refreshTokenSecret)
}
