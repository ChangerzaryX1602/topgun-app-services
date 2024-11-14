package prediction

type predictionService struct {
	predictionRepository PredictionRepository
}

func NewPredictionService(predictionRepository PredictionRepository) PredictionService {
	return &predictionService{predictionRepository}
}
func (s *predictionService) CreatePrediction(data Prediction) error {
	return s.predictionRepository.CreatePrediction(data)
}
