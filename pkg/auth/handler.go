package auth

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	helpers "github.com/zercle/gofiber-helpers"

	"top-gun-app-services/pkg/models"
)

type AuthHandler struct {
	authService AuthService
	jwt         models.JwtResources
}

func NewAuthHandler(authRoute fiber.Router, auth AuthService, jwt models.JwtResources) {
	handler := &AuthHandler{
		authService: auth,
		jwt:         jwt,
	}
	authRoute.Post("/", handler.Login())
}

// @Summary Login
// @Description get code from https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=2006502412&redirect_uri=https://app.finizer.co/callback&state=random&scope=profile%20openid%20email&nonce=finizer
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body Oauth true "Oauth"
// @Router /api/v1/auth/ [post]
func (h *AuthHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		request := Oauth{}
		err = c.BodyParser(&request)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    http.StatusBadRequest,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(http.StatusBadRequest).JSON(responseForm)
		}
		login, err := h.authService.Login(request.Code)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    http.StatusUnauthorized,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(http.StatusUnauthorized).JSON(responseForm)
		}
		token := jwt.NewWithClaims(h.jwt.JwtSigningMethod, &jwt.RegisteredClaims{})
		claims := token.Claims.(*jwt.RegisteredClaims)
		claims.Subject = strconv.Itoa(login.ID)
		claims.Issuer = c.Hostname()
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
		signToken, err := token.SignedString(h.jwt.JwtSignKey)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    http.StatusUnauthorized,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(http.StatusUnauthorized).JSON(responseForm)
		}
		responseForm.Result = fiber.Map{
			"token": signToken,
			"user":  login,
		}
		responseForm.Messages = []string{login.Username + " have been login successfully."}
		responseForm.Success = true
		return c.Status(http.StatusOK).JSON(responseForm)
	}
}

//
