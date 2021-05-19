package mq

import (
	"github.com/VlasovArtem/distributed-system-example/authors/internal/config"
	"github.com/streadway/amqp"
	"log"
)

type MessageQueueService struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func (mqs *MessageQueueService) CloseMessageQueueService() {
	mqs.channel.Close()
	mqs.connection.Close()
}

func StartMessageQueueConnection(config *config.Config) *MessageQueueService {
	connection, err := amqp.Dial(config.MQ.URL)
	if err != nil {
		log.Fatal("Could not connect to a Message Queue server", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal("Could not open channel for a MQ", err)
	}

	return &MessageQueueService{
		connection: connection,
		channel: channel,
	}
}
