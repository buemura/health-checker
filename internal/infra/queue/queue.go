package queue

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/buemura/health-checker/config"
	"github.com/buemura/health-checker/internal/core/dto"
	"github.com/buemura/health-checker/internal/core/usecase"
	"github.com/buemura/health-checker/internal/infra/database"
	amqp "github.com/rabbitmq/amqp091-go"
)

const NOTIFY_ENDPOINT_DOWN_QUEUE = "notify.endpoint.down"

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func StartConsume() {
	// Connect to rabbitmq broker
	conn, err := amqp.Dial(config.BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare/Create Queue
	_, err = ch.QueueDeclare(
		NOTIFY_ENDPOINT_DOWN_QUEUE, // name
		false,                      // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Consume Queue
	msgs, err := ch.Consume(
		NOTIFY_ENDPOINT_DOWN_QUEUE, // queue
		"",                         // consumer
		true,                       // auto-ack
		false,                      // exclusive
		false,                      // no-local
		false,                      // no-wait
		nil,                        // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Consumed messagem from queue: notify.endpoint.down")

			switch d.RoutingKey {
			case NOTIFY_ENDPOINT_DOWN_QUEUE:
				var in *dto.CreateEndpointIn
				err := json.Unmarshal([]byte(d.Body), &in)
				if err != nil {
					log.Fatalf(err.Error())
				}

				er := database.NewEndpointRepositoryImpl(database.DB)
				uc := usecase.NewCreateEndpoint(er)
				uc.Execute(in)
			}
		}
	}()

	fmt.Println("⇨ RabbitMQ Consumer started")
	<-forever
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, body string, exName string) error {
	err := ch.Publish(
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
