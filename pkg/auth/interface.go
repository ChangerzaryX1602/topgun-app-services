package auth

import "top-gun-app-services/pkg/user"

type AuthRepository interface {
	Login(user.User) (*user.User, error)
	Register(user.User) (*user.User, error)
}

type AuthService interface {
	Register(user.User) (*UUID, error)
	Login(LoginBody) (*LoginBody, error)
}
