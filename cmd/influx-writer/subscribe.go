package main

import (
	"github.com/streadway/amqp"
	"log"
	"wcrwg-iot-ingress/pkg/utils"
)

func SubscribeThread() {
	log.Println("Connecting to Rabbitmq")

	conn, err := amqp.Dial("amqp://" + myConfiguration.AmqpUser + ":" + myConfiguration.AmqpPassword + "@" + myConfiguration.AmqpHost + ":" + myConfiguration.AmqpPort + "/")
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
	utils.FailOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		myConfiguration.AmqpQueue, // name
		true,                      // durable
		false,                     // delete when usused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		100,   // prefetch count
		0,     // prefetch size
		false, // global
	)
	utils.FailOnError(err, "Failed to set queue QoS")

	err = ch.QueueBind(
		q.Name,                       // queue name
		"",                           // routing key
		myConfiguration.AmqpExchange, // exchange
		false,
		nil)
	utils.FailOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	log.Printf(" [a] Waiting for messages. To exit press CTRL+C")

	for d := range msgs {
		log.Printf(" [a] Message received %s", d.Body)
		subscribeChannel <- d
	}

	log.Fatal("No more messages on subscribe channel")
}
