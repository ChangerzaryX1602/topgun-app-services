package prediction

type PredictionRepository interface {
	CreatePrediction(Prediction) error
}
type PredictionService interface {
	CreatePrediction(Prediction) error
}
