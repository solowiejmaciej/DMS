package services

import mqtt "github.com/eclipse/paho.mqtt.golang"

func PublishMessageToTopic(mqqClient mqtt.Client, topic string, message string) {
	token := mqqClient.Publish(topic, 0, false, message)
	token.Wait()
}
