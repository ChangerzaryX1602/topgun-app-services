package auth

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type authRepository struct {
	Resource *gorm.DB
}

func NewAuthRepository(resources *gorm.DB) AuthRepository {
	return &authRepository{
		Resource: resources,
	}
}
func (r *authRepository) Login(email string, users User) (*User, error) {
	if r.Resource == nil {
		err := fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
		return nil, err
	}
	var existingUser User
	err := r.Resource.Preload(clause.Associations).First(&existingUser, "email = ?", email).Error
	if err != nil {
		users.CreatedAt = time.Now()
		err := r.Resource.Preload(clause.Associations).Create(&users).Error
		if err != nil {
			return nil, err
		}
		return &users, err
	}
	users.UpdatedAt = time.Now()
	err = r.Resource.Model(&existingUser).Preload(clause.Associations).Updates(&users).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("existingUser:\n", existingUser)
	return &existingUser, nil
}
