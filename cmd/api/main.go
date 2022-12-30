package main

import (
	"encoding/json"
	"fmt"
	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"wcrwg-iot-ingress/cmd/api/tts"
	"wcrwg-iot-ingress/pkg/types"
	"wcrwg-iot-ingress/pkg/utils"
)

var publishChannel = make(chan types.IotMessage, 100)

type Configuration struct {
	AmqpHost     string `envconfig:"AMQP_HOST"`
	AmqpPort     string `envconfig:"AMQP_PORT"`
	AmqpUser     string `envconfig:"AMQP_USER"`
	AmqpPassword string `envconfig:"AMQP_PASSWORD"`
	AmqpExchange string `envconfig:"AMQP_EXCHANGE"`

	HttpListenAddress string `envconfig:"HTTP_LISTEN_ADDRESS"`
}

var myConfiguration = Configuration{
	AmqpHost:     "localhost",
	AmqpPort:     "5672",
	AmqpUser:     "guest",
	AmqpPassword: "guest",
	AmqpExchange: "new-iot-messages",

	HttpListenAddress: ":8080",
}

func main() {
	err := envconfig.Process("", &myConfiguration)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("[Configuration]\n%s\n", utils.PrettyPrint(myConfiguration))

	router := Routes(publishChannel)

	log.Print("[Routes]")
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	// Start the thread that process the queue
	go publishFromChannel()

	// Start the http endpoint
	log.Fatal(http.ListenAndServe(myConfiguration.HttpListenAddress, router))
}

func Routes(publishChannel chan types.IotMessage) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RealIP,
		//middleware.Logger,
		middleware.Compress(5),
		middleware.StripSlashes,
		middleware.Recoverer,
		chiprometheus.NewMiddleware("iot-ingress-api", 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1, 1.5, 2, 5, 10, 100, 1000, 10000),
	)

	router.Handle("/metrics", promhttp.Handler())

	router.Mount("/tts", tts.Routes(publishChannel))
	//router.Mount("/sigfox", sigfox.Routes(publishChannel))

	return router
}

func publishFromChannel() {
	conn, err := amqp.Dial("amqp://" + myConfiguration.AmqpUser + ":" + myConfiguration.AmqpPassword + "@" + myConfiguration.AmqpHost + ":" + myConfiguration.AmqpPort + "/")
	//if err != nil {
	//	log.Print("Error connecting to RabbitMQ")
	//	return
	//}
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		myConfiguration.AmqpExchange, // name
		"fanout",                     // type
		true,                         // durable
		false,                        // auto-deleted
		false,                        // internal
		false,                        // no-wait
		nil,                          // arguments
	)

	//var message map[string]interface{}
	for {
		message := <-publishChannel
		log.Printf("Publishing message")

		data, err := json.Marshal(message)
		if err != nil {
			fmt.Printf("marshal failed: %s", err)
			continue
		}

		err = ch.Publish(
			myConfiguration.AmqpExchange, // exchange
			"",                           // routing key
			false,                        // mandatory
			false,                        // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(data),
			})
		utils.FailOnError(err, "Failed to publish a message")
	}
}
