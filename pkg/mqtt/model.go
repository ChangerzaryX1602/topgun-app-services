package mqtt

import "time"

type MQTT struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Message   Predict   `json:"predict" gorm:"type:jsonb"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}
type MQTTRequest struct {
	Message string `json:"message"`
	Topic   string `json:"topic"`
}
type Predict struct {
	Predict       string  `json:"predict"`
	VoiceRealTime float64 `json:"voice_real_time"`
}
