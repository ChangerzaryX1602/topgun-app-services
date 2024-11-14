package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MqttService interface {
	MessagePubHandler(client mqtt.Client, msg mqtt.Message)
	PublishMessage(topic string, message []byte) error
}
type MqttRepository interface {
	CreateMessage(MQTT) (MQTT, error)
}
