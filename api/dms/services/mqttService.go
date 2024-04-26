package services

import (
	"dms/entites"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
)

var MqttClients = make(map[string]mqtt.Client)

func ConfigureMqttForDevice(deviceConfiguration entites.DeviceConfiguration) {
	go func() {
		log.Info("Connecting to MQTT broker for device ", deviceConfiguration.DeviceId)
		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tls://%s:%d", deviceConfiguration.MqttBrokerHost, deviceConfiguration.MqttBrokerPort))
		opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
		opts.SetUsername(deviceConfiguration.MqttUsername)
		opts.SetPassword(deviceConfiguration.MqttPassword)

		MqttClient := mqtt.NewClient(opts)
		if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
			log.Error(fmt.Sprintf("Error connecting to MQTT broker for device %s: %s", deviceConfiguration.DeviceId, token.Error()))
			return
		}
		log.Info("Connection to MQTT broker for device ", deviceConfiguration.DeviceId, " established")
		MqttClient.Connect()
		MqttClients[deviceConfiguration.DeviceId] = MqttClient
	}()
}
func GetMqttClient(deviceId string) mqtt.Client {
	return MqttClients[deviceId]
}

func PublishMessageToTopic(deviceId string, topic string, message string) {
	mqqtClient := GetMqttClient(deviceId)
	token := mqqtClient.Publish(topic, 1, false, message)
	token.Wait()
}

func TestMqttConnection(host string, port int, username string, password string) (bool, error) {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tls://%s:%d", host, port))
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.SetUsername(username)
	opts.SetPassword(password)

	MqttClient := mqtt.NewClient(opts)
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		return false, token.Error()
	}
	MqttClient.Disconnect(0)
	return true, nil
}
