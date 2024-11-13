package workshop

import (
	"top-gun-app-services/pkg/models"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

type workshopHandler struct {
	service WorkshopService
}

func NewWorkshopHandler(route fiber.Router, service WorkshopService) {
	handler := &workshopHandler{service}
	//CRUD
	route.Post("/", handler.CreateMachine())
	route.Get("/", handler.GetMachines())
	route.Get("/:id", handler.GetMachine())
	route.Put("/:id", handler.UpdateMachine())
	route.Delete("/:id", handler.DeleteMachine())
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
		paginate := models.Paginate{}
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
		return c.Status(fiber.StatusNoContent).JSON(nil)
	}
}
