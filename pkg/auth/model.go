package auth

import (
	"top-gun-app-services/pkg/user"

	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type Oauth struct {
	Code string `json:"code"`
}
type AuthResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Picture  string `json:"picture"`
	Sub      string `json:"sub"`
}
type User user.User
