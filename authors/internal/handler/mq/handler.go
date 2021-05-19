package mq

import (
	"encoding/json"
	"github.com/VlasovArtem/distributed-system-example/authors/internal/service"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func (mqs *MessageQueueService) StartMessageQueueConsumer(s *service.Service) {
	messages, err := mqs.channel.Consume(
		service.QueueName, // queue name
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no local
		false,             // no wait
		nil,               // arguments
	)
	if err != nil {
		log.Fatal("Could not consume from channel", err)
	}

	go consumeMessages(s, messages)
}

func consumeMessages(s *service.Service, messages <-chan amqp.Delivery) {
	for message := range messages {
		bookAndAuthor := service.BookAndAuthor{}

		err := json.Unmarshal(message.Body, &bookAndAuthor)

		if err != nil {
			log.Println("Could not parse message")
		} else {
			s.AddAuthor(
				service.Author{
					ID:        strconv.Itoa(bookAndAuthor.AuthorID),
					FirstName: bookAndAuthor.FirstName,
					LastName:  bookAndAuthor.LastName,
				},
			)
		}
	}
}
