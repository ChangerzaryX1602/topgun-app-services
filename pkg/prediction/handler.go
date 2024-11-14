package prediction

import (
	"top-gun-app-services/internal/handlers"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

type predictionHandler struct {
	predictionService PredictionService
	auth              *handlers.RouterResources
}

func NewPredictionHandler(route fiber.Router, predictionService PredictionService, router *handlers.RouterResources) {
	handler := &predictionHandler{predictionService: predictionService, auth: router}
	route.Post("/", handler.auth.ReqAuthHandler(0), handler.CreatePrediction())
}

// @Summary Create Prediction
// @Description Create Prediction
// @Tags Integration
// @Accept json
// @Produce json
// @Param body body Prediction true "Prediction"
// @Success 201 {object} Prediction
// @Security ApiKeyAuth
// @Router /prediction/ [post]
func (h *predictionHandler) CreatePrediction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data Prediction
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
		err := h.predictionService.CreatePrediction(data)
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
			Data:    data,
		})
	}
}
