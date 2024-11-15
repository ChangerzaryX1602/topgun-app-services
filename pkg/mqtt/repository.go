package mqtt

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type mqttRepository struct {
	db *gorm.DB
}

func NewMQTTRepository(db *gorm.DB) MqttRepository {
	return &mqttRepository{db: db}
}
func (m Predict) Value() (driver.Value, error) {
	return json.Marshal(m)
}
func (m *Predict) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid type")
	}
	return json.Unmarshal(bytes, m)
}
func (r *mqttRepository) CreatePredict(mqttRequest PredictData) (PredictData, error) {
	if r.db == nil {
		return mqttRequest, fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
	}
	err := r.db.Create(&mqttRequest).Error
	if err != nil {
		return mqttRequest, err
	}
	fmt.Println("Message created: ", mqttRequest)
	return mqttRequest, nil
}
func (r *mqttRepository) CreateVoice(mqttRequest VoiceData) (VoiceData, error) {
	if r.db == nil {
		return mqttRequest, fiber.NewError(fiber.StatusServiceUnavailable, "Database server has gone away")
	}
	err := r.db.Create(&mqttRequest).Error
	if err != nil {
		return mqttRequest, err
	}
	fmt.Println("Message created: ", mqttRequest)
	return mqttRequest, nil
}
