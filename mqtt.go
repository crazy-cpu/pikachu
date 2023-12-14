package pikachu

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var (
	Conn mqtt.Client
)

type (
	Client  mqtt.Client
	Message mqtt.Message
	Handler mqtt.MessageHandler
)

func initMqtt(broker string) {
	ops := mqtt.NewClientOptions()
	ops.AutoReconnect = true
	ops.ConnectTimeout = 5 * time.Second
	ops.AddBroker(broker)

	Conn = mqtt.NewClient(ops)
	if token := Conn.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return
}

func Publish(topic string, payload any) error {
	token := Conn.Publish(topic, 0, false, payload)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
