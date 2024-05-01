package services

import (
	"dms/entities"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var MqttClients = make(map[string]mqtt.Client)

func ConfigureMqttForDevice(deviceConfiguration entities.DeviceConfiguration) {
	log.Info("Connecting to MQTT broker for device ", deviceConfiguration.DeviceId)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", deviceConfiguration.MqttBrokerHost, deviceConfiguration.MqttBrokerPort))
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID") + deviceConfiguration.DeviceId)
	opts.SetUsername(deviceConfiguration.MqttUsername)
	opts.SetPassword(deviceConfiguration.MqttPassword)

	MqttClient := mqtt.NewClient(opts)
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Error(fmt.Sprintf("Error connecting to MQTT broker for device %s: %s", deviceConfiguration.DeviceId, token.Error()))
		return
	}

	log.Info("Connection to MQTT broker for device ", deviceConfiguration.DeviceId, " established")
	MqttClients[deviceConfiguration.DeviceId] = MqttClient

	subscribeToAlive(deviceConfiguration.DeviceId)

	go KeepMqttClientConnected(deviceConfiguration)
}

func KeepMqttClientConnected(deviceConfiguration entities.DeviceConfiguration) {
	deviceId := deviceConfiguration.DeviceId
	for {
		MqttClient := MqttClients[deviceId]
		if !MqttClient.IsConnected() {
			log.Info("MQTT client disconnected, attempting to reconnect")
			if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
				log.Error(fmt.Sprintf("Error reconnecting to MQTT broker for device %s: %s", deviceId, token.Error()))
			} else {
				log.Info("Reconnected to MQTT broker for device ", deviceId)
				subscribeToAlive(deviceId)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func subscribeToAlive(deviceId string) {
	Client := MqttClients[deviceId]
	if token := Client.Subscribe("/api/"+deviceId+"/alive", 1, func(client mqtt.Client, message mqtt.Message) {
		log.Info("Received message for device ", deviceId)
		UpdateLastMessageTime(deviceId)
	}); token.Wait() && token.Error() != nil {
		log.Error(fmt.Sprintf("Error subscribing to topic for device %s: %s", deviceId, token.Error()))
	}

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
