package mqtt

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

type mqttHandler struct {
	mqttService MqttService
	mqtt        mqtt.Client
	mqttOption  *mqtt.ClientOptions
}

func NewMQttHandler(router fiber.Router, mqttService MqttService, mqtt mqtt.Client, mqttOption *mqtt.ClientOptions) {
	mqttHandler := mqttHandler{mqttService: mqttService, mqtt: mqtt, mqttOption: mqttOption}
	router.Post("/", mqttHandler.PostMqtt())
	mqttHandler.MqttSubscibeHandler(mqtt, mqttOption)
}
func (h mqttHandler) MqttSubscibeHandler(mqtt mqtt.Client, mqttOption *mqtt.ClientOptions) {
	if token := mqtt.Subscribe("arduino/temperature", 0, h.mqttService.MessagePubHandler); token.Wait() && token.Error() != nil {
		fmt.Printf("Error subscribing to topic: %v\n", token.Error())
	}
}
func (h mqttHandler) PostMqtt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req MQTTRequest
		err := c.BodyParser(&req)
		if err != nil {
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
		mqttResponse := helpers.ResponseForm{
			Success: true,
			Data:    req,
		}
		payload, err := json.Marshal(mqttResponse)
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
		token := h.mqtt.Publish("arduino/temperature", 0, false, payload)
		token.Wait()
		if token.Error() != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: token.Error().Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusOK).JSON(helpers.ResponseForm{
			Success:  true,
			Data:     req,
			Messages: []string{"Success"},
		})
	}
}
