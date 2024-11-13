package workshop

import "top-gun-app-services/pkg/models"

type WorkshopService interface {
	ConnectWebSocket(wsURL string, apiKey string)
	CreateMachine(RawData) (RawData, error)
	GetMachines(models.Paginate) ([]RawData, error)
	GetMachine(id string) (RawData, error)
	UpdateMachine(id string, data RawData) (RawData, error)
	DeleteMachine(id string) error
}
type WorkshopRepository interface {
	ProcessMessage(message []byte)
	CreateMachine(RawData) (RawData, error)
	GetMachines(models.Paginate) ([]RawData, error)
	GetMachine(id string) (RawData, error)
	UpdateMachine(id string, data RawData) (RawData, error)
	DeleteMachine(id string) error
}
