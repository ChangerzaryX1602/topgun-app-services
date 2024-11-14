package prediction

import "gorm.io/gorm"

type predictionRepository struct {
	db *gorm.DB
}

func NewPredictionRepository(db *gorm.DB) PredictionRepository {
	return &predictionRepository{db}
}
func (r predictionRepository) CreatePrediction(data Prediction) error {
	err := r.db.Create(&data)
	if err != nil {
		return err.Error
	}
	return nil
}
