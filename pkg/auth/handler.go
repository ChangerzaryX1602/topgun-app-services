package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	helpers "github.com/zercle/gofiber-helpers"
	"golang.org/x/crypto/bcrypt"

	"top-gun-app-services/pkg/models"
	"top-gun-app-services/pkg/user"
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
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body LoginBody true "Login Body"
// @Router /api/v1/auth/ [post]
func (h *AuthHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request LoginBody
		responseForm := helpers.ResponseForm{}
		// Parse the request body
		if err := c.BodyParser(&request); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"Error": "Failed to parse request body",
			})
		}

		// Validate the input credentials
		if request.Identifier == "" || request.Password == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"Error": "Invalid login credentials (Identifier or Password)",
			})
		}

		// Attempt to log in via the auth service
		response, err := h.authService.Login(request)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"Error": "Invalid identifier or password",
			})
		}

		// Compare the provided password with the stored hash
		err = bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(request.Password))
		if err != nil {
			loginLog := fmt.Sprintf("User %s failed to log in: incorrect password", response.UUID.String())
			fmt.Println("loginLog:", loginLog)
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"Error": "Incorrect password",
			})
		}

		token := jwt.NewWithClaims(h.jwt.JwtSigningMethod, &jwt.RegisteredClaims{})
		claims := token.Claims.(*jwt.RegisteredClaims)
		claims.Subject = response.UUID.String()
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

		// Log successful login
		loginLog := fmt.Sprintf("User %s has successfully logged in", response.UUID.String())
		fmt.Println("loginLog:", loginLog)

		// Return the generated token to the client
		return c.JSON(fiber.Map{"token": signToken})
	}
}

// @Summary Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body user.User true "User Body"
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		request := user.User{}
		err = c.BodyParser(&request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"Error": "Body Parser",
			})
		}

		response, err := h.authService.Register(request)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"Error": "Invalid Signup Credentials",
			})
		}
		registerLog := "User " + response.UUID.String() + " has been registered"
		return c.JSON(registerLog)

	}
}
