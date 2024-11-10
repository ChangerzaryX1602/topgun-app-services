package auth

type AuthRepository interface {
	Login(string, User) (*User, error)
}

type AuthService interface {
	Login(string) (*AuthResponse, error)
}
