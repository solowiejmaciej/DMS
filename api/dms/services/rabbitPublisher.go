package services

import (
	"dms/events"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func PublishSendConfirmPhoneNumber(userId uint, channel *amqp.Channel, client *redis.Client) {
	var generatedCode, err = GenerateCodeForUser(userId, client)
	if err != nil {
		log.Error("Error while generating code for user", err)
		return
	}
	event := events.ConfirmPhoneNumber{
		EventId: uuid.New(),
		UserId:  userId,
		Code:    generatedCode,
	}
	var messageBody, parsingError = json.Marshal(event)
	if parsingError != nil {
		log.Error("Error while parsing user to json", parsingError)
		return
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        messageBody,
	}
	publishError := channel.Publish(
		"",                  // exchange
		"dms-notifications", // queue name
		false,               // mandatory
		false,               // immediate
		message,             // message to publish
	)
	if publishError != nil {
		log.Error("Error while publishing message to the queue", publishError)
		return
	}
	log.Info("Event: ", event.EventId, " User: ", event.UserId)
}
