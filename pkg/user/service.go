package user

import (
	"time"

	"top-gun-app-services/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type userService struct {
	userRepository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}
func (s *userService) GetAllUsers(paginate models.Paginate) ([]User, models.Paginate, error) {
	return s.userRepository.GetAllUsers(paginate)
}
func (s *userService) GetMe(id string) (User, error) {
	user, err := s.userRepository.GetUserByID(id)
	return *user, err
}
func (s *userService) GetUser(id string) (User, error) {
	user, err := s.userRepository.GetUserByID(id)
	return *user, err
}
func (s *userService) DeleteByID(id string) error {
	return s.userRepository.DeleteByID(id)
}
func (s *userService) UpdateMe(id string, user User) error {
	updatedData := User{
		Name:      user.Name,
		Email:     user.Email,
		Birthdate: user.Birthdate,
		Gender:    user.Gender,
		UpdatedAt: time.Now(),
	}
	return s.userRepository.UpdateByID(id, updatedData)
}
func (s *userService) UpdateByID(id string, user User) error {
	updatedData := User{
		Name:      user.Name,
		Email:     user.Email,
		Birthdate: user.Birthdate,
		Gender:    user.Gender,
		UpdatedAt: time.Now(),
	}
	return s.userRepository.UpdateByID(id, updatedData)
}
func (s *userService) SearchUser(keyword string) ([]User, error) {
	users, err := s.userRepository.SearchUser(keyword)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	return users, nil
}
