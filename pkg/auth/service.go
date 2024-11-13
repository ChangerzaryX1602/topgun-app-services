package auth

import (
	"strings"
	"time"
	"top-gun-app-services/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepository AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{
		authRepository: repo,
	}
}
func (s *authService) Register(req user.User) (*UUID, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if req.Username == "" || req.Password == "" || req.Email == "" {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	user := user.User{
		ID:        uuid.New(),
		NameEn:    req.NameEn,
		Username:  req.Username,
		Password:  string(password),
		Email:     req.Email,
		Language:  "th",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newUser, err := s.authRepository.Register(user)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to register user")
	}

	userResponse := UUID{
		UUID: newUser.ID,
	}
	return &userResponse, nil
}
func (s *authService) Login(req LoginBody) (*LoginBody, error) {
	user := user.User{
		Password: req.Password,
	}
	if strings.Contains(req.Identifier, "@") {
		user.Email = req.Identifier
	} else {
		user.Username = req.Identifier
	}

	loginUser, err := s.authRepository.Login(user)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to login")
	}

	allUserResponse := LoginBody{
		UUID:     loginUser.ID,
		Password: loginUser.Password,
	}
	return &allUserResponse, nil
}
