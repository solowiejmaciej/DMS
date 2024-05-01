package initializers

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"os"
)

var ChannelRabbitMQ *amqp.Channel

func ConnectToRabbitMq() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	ChannelRabbitMQ, err = connectRabbitMQ.Channel()

	_, err = ChannelRabbitMQ.QueueDeclare(
		"dms-notifications", // queue name
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	log.Info("Successfully connected to RabbitMQ")
}
