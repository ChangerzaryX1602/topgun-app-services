package user

import (
	"top-gun-app-services/pkg/models"
)

type UserRepository interface {
	GetAllUsers(models.Paginate) ([]User, models.Paginate, error)
	GetUserByID(string) (*User, error)
	DeleteByID(string) error
	UpdateByID(string, User) error
	SearchUser(string) ([]User, error)
}

type UserService interface {
	GetAllUsers(models.Paginate) ([]User, models.Paginate, error)
	GetMe(string) (User, error)
	GetUser(string) (User, error)
	DeleteByID(string) error
	UpdateMe(string, User) error
	UpdateByID(string, User) error
	SearchUser(string) ([]User, error)
}
