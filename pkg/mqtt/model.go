package mqtt

import "time"

type VoiceData struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Voice     Voice     `json:"voice" gorm:"type:jsonb"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}
type PredictData struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Predict   Predict   `json:"predict" gorm:"type:jsonb"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}
type MQTTRequest struct {
	Message string `json:"message"`
	Topic   string `json:"topic"`
}
type Predict struct {
	Predict string `json:"predict"`
	Time    string `json:"time"`
}
type Voice struct {
	VoiceRealTime float64 `json:"voice_real_time"`
}
