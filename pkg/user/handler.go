package user

import (
	"fmt"

	"top-gun-app-services/internal/handlers"
	"top-gun-app-services/pkg/models"
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

type UserHandler struct {
	userService UserService
	auth        *handlers.RouterResources
}

func NewUserHandler(userRoute fiber.Router, us UserService, router *handlers.RouterResources) {
	handler := &UserHandler{
		userService: us,
		auth:        router,
	}
	userRoute.Get("/search", handler.auth.ReqAuthHandler(0), handler.SearchUser())
	userRoute.Get("/", handler.auth.ReqAuthHandler(0), handler.GetAllUsers())
	userRoute.Get("/me", handler.auth.ReqAuthHandler(0), handler.GetMe())
	userRoute.Get("/:id", handler.auth.ReqAuthHandler(0), handler.GetUser())
	userRoute.Delete("/:id", handler.auth.ReqAuthHandler(0), handler.DeleteByID())
	userRoute.Put("/me", handler.auth.ReqAuthHandler(0), handler.UpdateMe())
	userRoute.Put("/:id", handler.auth.ReqAuthHandler(0), handler.UpdateByID())
}

// @Summary Get All Users
// @Description Get All Users
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/v1/users/ [get]
func (h *UserHandler) GetAllUsers() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		userID := c.Locals("user_id").(string)
		req := models.Paginate{}
		err = c.QueryParser(&req)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusBadRequest,
					Message: err.Error(),
					Source:  helpers.WhereAmI(),
				},
			}
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}
		users, paginate, err := h.userService.GetAllUsers(req)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusInternalServerError,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}
		responseForm.Result = fiber.Map{
			"users":    users,
			"paginate": paginate,
		}
		responseForm.Messages = []string{fmt.Sprintf("user id %v have been get all users data successfully", userID)}
		responseForm.Success = true
		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}

// @Summary Get Me
// @Description Get Me
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/v1/users/me [get]
func (h *UserHandler) GetMe() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		userID := c.Locals("user_id").(string)
		user, err := h.userService.GetMe(userID)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusInternalServerError,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}
		responseForm.Result = fiber.Map{
			"user": user,
		}
		responseForm.Messages = []string{fmt.Sprintf("user id %v have been get user %v data successfully", userID, userID)}
		responseForm.Success = true
		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}

// @Summary Get User
// @Description Get User
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		userID := c.Locals("user_id").(string)
		user, err := h.userService.GetUser(c.Params("id"))
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusInternalServerError,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}
		responseForm.Messages = []string{fmt.Sprintf("user id %v have been get user %v data successfully", userID, userID)}
		responseForm.Result = fiber.Map{
			"user": user,
		}
		responseForm.Success = true
		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}

// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteByID() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		userID := c.Locals("user_id").(string)
		err = h.userService.DeleteByID(c.Params("id"))
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusInternalServerError,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}
		responseForm.Messages = []string{fmt.Sprintf("user id %v have been delete user %v data successfully", userID, userID)}
		responseForm.Success = true
		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}

// @Summary Update Me
// @Description Update Me
// @Param user body User false "User Data"
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/v1/users/me [put]
func (h *UserHandler) UpdateMe() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		request := User{}
		userID := c.Locals("user_id").(string)
		err = c.BodyParser(&request)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusBadRequest,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}
		err = h.userService.UpdateMe(userID, request)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusInternalServerError,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}
		responseForm.Messages = []string{fmt.Sprintf("user id %v have been update user %v data successfully", userID, userID)}
		responseForm.Success = true
		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}

// @Summary Update User
// @Description Update User
// @Param id path string true "User ID"
// @Param user body User true "User Data"
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateByID() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		request := User{}
		userID := c.Locals("user_id").(string)
		err = c.BodyParser(&request)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusBadRequest,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusBadRequest).JSON(responseForm)
		}
		err = h.userService.UpdateByID(c.Params("id"), request)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusInternalServerError,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}
		responseForm.Messages = []string{fmt.Sprintf("user id %v have been update user %v data successfully", userID, userID)}
		responseForm.Success = true
		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}

// @Summary Search User
// @Description Search User
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param keyword query string true "Keyword"
// @Router /api/v1/users/search [get]
func (h *UserHandler) SearchUser() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		responseForm := helpers.ResponseForm{}
		keyword := c.Query("keyword")
		res, err := h.userService.SearchUser(keyword)
		if err != nil {
			responseForm.Errors = []helpers.ResponseError{
				{
					Code:    fiber.StatusInternalServerError,
					Source:  helpers.WhereAmI(),
					Message: err.Error(),
				},
			}
			return c.Status(fiber.StatusInternalServerError).JSON(responseForm)
		}
		responseForm.Result = res
		responseForm.Success = true
		return c.Status(fiber.StatusOK).JSON(responseForm)
	}
}
