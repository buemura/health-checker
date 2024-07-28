package main

import (
	"encoding/json"
	"log"

	"github.com/buemura/health-checker/config"
	"github.com/buemura/health-checker/internal/core/dto"
	"github.com/buemura/health-checker/internal/infra/database"
	"github.com/buemura/health-checker/internal/infra/event"
	"github.com/buemura/health-checker/internal/infra/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial(config.BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue.DeclareQueue(ch, queue.NOTIFY_ENDPOINT_DOWN_QUEUE)
	queue.DeclareQueue(ch, queue.NOTIFY_ENDPOINT_DOWN_DLQ)

	msgs := make(chan amqp.Delivery)
	go queue.Consume(ch, msgs, queue.NOTIFY_ENDPOINT_DOWN_QUEUE)

	for msg := range msgs {
		log.Printf("Consumed messagem from queue: notify.endpoint.down")

		switch msg.RoutingKey {
		case queue.NOTIFY_ENDPOINT_DOWN_QUEUE:
			var in *dto.CreateNotificationIn
			err := json.Unmarshal([]byte(msg.Body), &in)
			if err != nil {
				log.Fatalf(err.Error())
			}

			notificationEvent := event.NewNotificationEvent()
			_, err = notificationEvent.SendNotification(in)
			if err != nil {
				log.Println(err)
				err = queue.PublishToQueue(ch, string(msg.Body), queue.NOTIFY_ENDPOINT_DOWN_DLQ)
				if err != nil {
					log.Fatalf("Failed to send message to DLQ queue: %s", err)
				}
			}
		}
	}
}
