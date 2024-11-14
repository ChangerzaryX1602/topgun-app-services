package prediction

type Prediction struct {
	ID               int    `gorm:"primaryKey;autoIncrement"`
	TimeStamp        string `json:"timeStamp"`
	PredictionResult string `json:"result"`
}
