package workshop

import (
	"strconv"
	"top-gun-app-services/internal/handlers"
	"top-gun-app-services/pkg/models"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

type workshopHandler struct {
	service WorkshopService
	auth    *handlers.RouterResources
}

func NewWorkshopHandler(route fiber.Router, service WorkshopService, router *handlers.RouterResources) {
	handler := &workshopHandler{service: service, auth: router}
	//CRUD
	route.Post("/", handler.auth.ReqAuthHandler(0), handler.CreateMachine())
	route.Get("/", handler.auth.ReqAuthHandler(0), handler.GetMachines())
	route.Get("/:id", handler.auth.ReqAuthHandler(0), handler.GetMachine())
	route.Put("/:id", handler.auth.ReqAuthHandler(0), handler.UpdateMachine())
	route.Delete("/:id", handler.auth.ReqAuthHandler(0), handler.DeleteMachine())
}
func (h *workshopHandler) CreateMachine() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data RawData
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		machine, err := h.service.CreateMachine(data)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusCreated).JSON(helpers.ResponseForm{
			Success: true,
			Data:    machine,
		})
	}
}
func (h *workshopHandler) GetMachines() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginate := models.DatePicker{}
		if err := c.QueryParser(&paginate); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		machines, err := h.service.GetMachines(paginate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusOK).JSON(helpers.ResponseForm{
			Success: true,
			Data:    machines,
		})
	}
}
func (h *workshopHandler) GetMachine() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		machine, err := h.service.GetMachine(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusOK).JSON(helpers.ResponseForm{
			Success: true,
			Data:    machine,
		})
	}
}
func (h *workshopHandler) UpdateMachine() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idString, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		var data RawData
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		machine, err := h.service.UpdateMachine(id, data)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		machine.ID = idString
		return c.Status(fiber.StatusOK).JSON(helpers.ResponseForm{
			Success: true,
			Data:    machine,
		})
	}
}
func (h *workshopHandler) DeleteMachine() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := h.service.DeleteMachine(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusNoContent).JSON(helpers.ResponseForm{
			Success:  true,
			Messages: []string{"Machine deleted successfully"},
		})
	}
}
