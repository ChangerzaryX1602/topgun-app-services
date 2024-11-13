package auth

import (
	"fmt"
	"top-gun-app-services/pkg/user"

	"gorm.io/gorm"
)

type authRepository struct {
	Resource *gorm.DB
}

func NewAuthRepository(resources *gorm.DB) AuthRepository {
	return &authRepository{
		Resource: resources,
	}
}
func (r *authRepository) Register(user user.User) (*user.User, error) {
	err := r.Resource.Create(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (r *authRepository) Login(user user.User) (*user.User, error) {
	if user.Username != "" {
		fmt.Println("username", user.Username)
		err := r.Resource.First(&user, "username = ?", user.Username)
		if err.Error != nil {
			return nil, err.Error
		}
	}
	if user.Email != "" {
		fmt.Println("email", user.Email)
		err := r.Resource.First(&user, "email = ?", user.Email)
		if err.Error != nil {
			return nil, err.Error
		}
	}

	return &user, nil
}
