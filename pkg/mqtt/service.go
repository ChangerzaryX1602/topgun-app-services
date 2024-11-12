package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqttService struct {
	mqttRepository MqttRepository
}

func NewMQttService(mqttRepository MqttRepository) MqttService {
	return &mqttService{mqttRepository: mqttRepository}
}
func (s mqttService) MessagePubHandler(client mqtt.Client, msg mqtt.Message) {
	mqttRequest := MQTT{
		Message:   string(msg.Payload()),
		Topic:     msg.Topic(),
		CreatedAt: time.Now(),
	}
	message, err := s.mqttRepository.CreateMessage(mqttRequest)
	if err != nil {
		fmt.Printf("Error creating message: %v\n", err)
	}
	fmt.Println("Message created: ", message)
}
