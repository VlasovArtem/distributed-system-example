package mq

import (
	"encoding/json"
	"fmt"
	"github.com/VlasovArtem/distributed-system-example/books/internal/config"
	"github.com/VlasovArtem/distributed-system-example/books/internal/model"
	"github.com/streadway/amqp"
	"log"
)

var QueueName = "BookAndAuthor"

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

func (mqs *MessageQueueService) DeclareQueue() *MessageQueueService {
	if _, err := mqs.channel.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Fatal(fmt.Sprintf("Could not declare queue %s", QueueName))
	}

	return mqs
}

func (mqs *MessageQueueService) PublishBookAndAuthorMessage(bookAndAuthor model.BookAndAuthor) error  {
	body, _ := json.Marshal(bookAndAuthor)

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}

	return mqs.channel.Publish(
		"",
		QueueName,
		false,
		false,
		message,
	)
}