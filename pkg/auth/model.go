package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type Oauth struct {
	Code string `json:"code"`
}
type LoginBody struct {
	UUID       uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	Password   string    `json:"password"`
}

type UUID struct {
	UUID uuid.UUID `json:"id"`
}
type LineInfo struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	PictureURL  string `json:"picture_url"`
}
type LineResponse struct {
	Iss     string   `json:"iss"`
	Sub     string   `json:"sub"`
	Aud     string   `json:"aud"`
	Exp     int64    `json:"exp"`
	Iat     int64    `json:"iat"`
	Nonce   string   `json:"nonce"`
	Amr     []string `json:"amr"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Email   string   `json:"email"`
}
