package mqtt

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type mqttRepository struct {
	db *gorm.DB
}

func NewMQTTRepository(db *gorm.DB) MqttRepository {
	return &mqttRepository{db: db}
}
func (r *mqttRepository) CreateMessage(mqttRequest MQTT) (MQTT, error) {
	if r.db == nil {
		return mqttRequest, fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
	}
	err := r.db.Create(&mqttRequest).Error
	if err != nil {
		return mqttRequest, err
	}
	return mqttRequest, nil
}
