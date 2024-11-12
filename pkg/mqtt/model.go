package mqtt

import "time"

type MQTT struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Message   string    `json:"message"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}
type MQTTRequest struct {
	Message string `json:"message"`
	Topic   string `json:"topic"`
}
