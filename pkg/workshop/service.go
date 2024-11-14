package workshop

import (
	"log"
	"top-gun-app-services/pkg/models"

	"github.com/gorilla/websocket"
)

type workshopService struct {
	repo WorkshopRepository
}

func NewWorkshopService(repo WorkshopRepository) WorkshopService {
	return &workshopService{repo}
}
func (s workshopService) ConnectWebSocket(wsURL string, apiKey string) {
	log.Printf("Connecting to WebSocket: %s", wsURL)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to WebSocket server.")

	// Send the API key upon connection
	err = conn.WriteMessage(websocket.TextMessage, []byte(apiKey))
	if err != nil {
		log.Fatalf("Failed to send API key: %v", err)
	}
	log.Println("API key sent.")
	// Receive data from the WebSocket and write to DB
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading WebSocket message: %v", err)
			break
		}
		s.repo.ProcessMessage(message)
	}
}
func (s workshopService) CreateMachine(data RawData) (RawData, error) {
	return s.repo.CreateMachine(data)
}
func (s workshopService) GetMachines(paginate models.DatePicker) ([]RawData, error) {
	return s.repo.GetMachines(paginate)
}
func (s workshopService) GetMachine(id string) (RawData, error) {
	return s.repo.GetMachine(id)
}
func (s workshopService) UpdateMachine(id string, data RawData) (RawData, error) {
	return s.repo.UpdateMachine(id, data)
}
func (s workshopService) DeleteMachine(id string) error {
	return s.repo.DeleteMachine(id)
}
