package user

import (
	"top-gun-app-services/pkg/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type userRepository struct {
	Resource *gorm.DB
}

func NewUserRepository(resources *gorm.DB) UserRepository {
	return &userRepository{
		Resource: resources,
	}
}

func (r *userRepository) GetAllUsers(paginate models.Paginate) ([]User, models.Paginate, error) {
	if r.Resource == nil {
		return nil, models.Paginate{}, fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
	}
	var users []User
	err := r.Resource.Find(&users).Offset(paginate.Offset).Limit(paginate.Limit).Error
	if err != nil {
		return nil, models.Paginate{}, err
	}
	err = r.Resource.Find(&users).Count(&paginate.Total).Error
	if err != nil {
		return nil, models.Paginate{}, err
	}
	paginateData := models.Paginate{
		Offset: paginate.Offset,
		Limit:  paginate.Limit,
		Total:  paginate.Total,
	}
	return users, paginateData, err
}
func (r *userRepository) GetUserByID(id string) (*User, error) {
	if r.Resource == nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
	}
	var user User
	err := r.Resource.First(&user, "id = ?", id).Error
	return &user, err
}
func (r *userRepository) DeleteByID(id string) error {
	if r.Resource == nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
	}
	err := r.Resource.Delete(&User{}, "id = ?", id).Error
	return err
}
func (r *userRepository) UpdateByID(id string, user User) error {
	if r.Resource == nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
	}
	err := r.Resource.Model(&User{}).Where("id = ?", id).Updates(user).Error
	return err
}
func (r *userRepository) SearchUser(keyword string) ([]User, error) {
	if r.Resource == nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
	}
	keyword = "%" + keyword + "%"
	var users []User
	err := r.Resource.Where("name LIKE ? OR email LIKE ?", keyword, keyword).Limit(5).Find(&users).Error
	return users, err
}
