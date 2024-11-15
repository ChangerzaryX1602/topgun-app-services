package datasources

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

func MqttConnect(broker string, clientID string) (mqtt.Client, *mqtt.ClientOptions, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.Username = viper.GetString("mqtt.username")
	opts.Password = viper.GetString("mqtt.password")
	opts.OnConnect = func(client mqtt.Client) {
		if token := client.Subscribe("arduino/temperature", 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}
	// opts.SetDefaultPublishHandler(mqttService.DefaultMessagePubHandler)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, nil, token.Error()
	}
	fmt.Println("Connected to MQTT broker")
	return client, opts, nil
}
