package database

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"rental-management-api/config"
	"rental-management-api/internal/constant"
	"rental-management-api/internal/entity"
)

const (
	seedAdminName  = "Administrator"
	seedAdminEmail = "admin@rental.local"
)

func SeedAdminUser(db *gorm.DB, cfg config.Config) error {
	var user entity.User
	err := db.Where("email = ?", seedAdminEmail).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("query admin user: %w", err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
		if hashErr != nil {
			return fmt.Errorf("hash admin password: %w", hashErr)
		}

		admin := entity.User{
			Name:     seedAdminName,
			Email:    seedAdminEmail,
			Role:     constant.UserRoleAdmin,
			Password: string(hashedPassword),
		}
		if createErr := db.Create(&admin).Error; createErr != nil {
			return fmt.Errorf("create admin user: %w", createErr)
		}
		return nil
	}

	if user.Role != constant.UserRoleAdmin {
		if updateErr := db.Model(&user).Update("role", constant.UserRoleAdmin).Error; updateErr != nil {
			return fmt.Errorf("update admin role: %w", updateErr)
		}
	}

	return nil
}
