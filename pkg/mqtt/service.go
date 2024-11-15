package mqtt

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqttService struct {
	mqttRepository MqttRepository
	mqtt           mqtt.Client
	mqttOption     *mqtt.ClientOptions
}

func NewMQttService(mqttRepository MqttRepository, mqtt mqtt.Client, mqttOption *mqtt.ClientOptions) MqttService {
	return &mqttService{mqttRepository: mqttRepository, mqtt: mqtt, mqttOption: mqttOption}
}
func (s mqttService) MessagePubHandler(client mqtt.Client, msg mqtt.Message) {
	var predict Predict
	err := json.Unmarshal(msg.Payload(), &predict)
	if err != nil {
		fmt.Printf("Error unmarshalling message: %v\n", err)
		return
	}
	mqttRequest := MQTT{
		Message:   predict,
		Topic:     msg.Topic(),
		CreatedAt: time.Now(),
	}
	message, err := s.mqttRepository.CreateMessage(mqttRequest)
	if err != nil {
		fmt.Printf("Error creating message: %v\n", err)
	}
	fmt.Println("Message created: ", message)
}
func (s mqttService) PublishMessage(topic string, message []byte) error {
	token := s.mqtt.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}
