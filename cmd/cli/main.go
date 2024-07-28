package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/buemura/health-checker/config"
	"github.com/buemura/health-checker/internal/core/dto"
	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/buemura/health-checker/internal/core/usecase"
	"github.com/buemura/health-checker/internal/infra/database"
	"github.com/buemura/health-checker/internal/infra/queue"
	"github.com/go-co-op/gocron/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ch *amqp.Channel
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	conn, err := amqp.Dial(config.BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = s.NewJob(
		gocron.DurationJob(5*time.Second),
		gocron.NewTask(func() {
			validateEndpoint()
		}),
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	s.Start()

	select {}
}

func validateEndpoint() {
	er := database.NewEndpointRepositoryImpl(database.DB)
	getEndpointListUC := usecase.NewGetEndpointList(er)
	updateEndpointUC := usecase.NewUpdateEndpoint(er)
	endpoints, err := getEndpointListUC.Execute()
	if err != nil {
		log.Fatalln(err.Error())
	}

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "URL", "Status", "Frequency (min)", "Last Checked", "Notify To"})

	for _, endpoint := range endpoints {
		RenderTable(t, endpoint)

		lastChecked := time.Now().Sub(endpoint.LastChecked.Add(3 * time.Hour)).Minutes()

		if int(lastChecked) > endpoint.CheckFrequency {
			log.Println("Checking endpoint:", endpoint.Url)

			response, err := http.Get(endpoint.Url)
			if err != nil || response.StatusCode != http.StatusOK {
				endpoint.Status = "DOWN"
				err := sendNotification(&dto.CreateNotificationIn{
					EndpointID:  endpoint.ID,
					Destination: endpoint.NotifyTo,
				})
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				endpoint.Status = "UP"
			}
			updateEndpointUC.Execute(endpoint)
		}
	}

	t.Render()
}

func RenderTable(t table.Writer, e *entity.Endpoint) {
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 4, Colors: text.Colors{text.FgYellow}},
	})
	t.AppendRow(table.Row{e.ID, e.Name, e.Url, e.Status, e.CheckFrequency, e.LastChecked.Format("2006-02-01 15:04:05"), e.NotifyTo})

}

func sendNotification(in *dto.CreateNotificationIn) error {
	payload, err := json.Marshal(in)
	if err != nil {
		return err
	}

	err = queue.PublishToQueue(ch, string(payload), queue.NOTIFY_ENDPOINT_DOWN_QUEUE)
	if err != nil {
		return err
	}

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
