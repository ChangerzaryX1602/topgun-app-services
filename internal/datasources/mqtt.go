package datasources

import (
	"fmt"
	"top-gun-app-services/pkg/utils"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MqttConnect(broker string, clientID string) (*mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected to MQTT broker")
		if token := client.Subscribe("arduino/temperature", 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}
	opts.SetDefaultPublishHandler(utils.MessagePubHandler)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &client, nil
}
