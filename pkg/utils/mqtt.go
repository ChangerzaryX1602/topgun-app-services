package utils

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	if msg.Topic() == "arduino/temperature" {
		fmt.Printf("Temperature: %s\n", msg.Payload())
	}
}
