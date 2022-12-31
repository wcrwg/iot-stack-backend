package main

import (
	"encoding/json"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/kelseyhightower/envconfig"
	"github.com/streadway/amqp"
	"log"
	"math"
	"net/http"
	"wcrwg-iot-ingress/pkg/influxdb"
	"wcrwg-iot-ingress/pkg/types"
	"wcrwg-iot-ingress/pkg/utils"
)

type Configuration struct {
	AmqpHost     string `envconfig:"AMQP_HOST"`
	AmqpPort     string `envconfig:"AMQP_PORT"`
	AmqpUser     string `envconfig:"AMQP_USER"`
	AmqpPassword string `envconfig:"AMQP_PASSWORD"`
	AmqpExchange string `envconfig:"AMQP_EXCHANGE"`
	AmqpQueue    string `envconfig:"AMQP_QUEUE"`

	InfluxdbServerUrl string `envconfig:"INFLUXDB_SERVER_URL"`
	InfluxdbAuthToken string `envconfig:"INFLUXDB_AUTH_TOKEN"`
	InfluxdbOrg       string `envconfig:"INFLUXDB_ORG"`
	InfluxdbBucket    string `envconfig:"INFLUXDB_BUCKET"`
}

var myConfiguration = Configuration{
	AmqpHost:     "localhost",
	AmqpPort:     "5672",
	AmqpUser:     "user",
	AmqpPassword: "password",
	AmqpExchange: "new-iot-messages",
	AmqpQueue:    "influxdb-insert",

	InfluxdbServerUrl: "http://localhost:8086",
	InfluxdbAuthToken: "token",
	InfluxdbBucket:    "bucket",
	InfluxdbOrg:       "org",
}

var subscribeChannel = make(chan amqp.Delivery)

func main() {
	err := envconfig.Process("", &myConfiguration)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("[Configuration]\n%s\n", utils.PrettyPrint(myConfiguration))

	influxdbCredentials := influxdb.InfluxCredentials{
		ServerUrl: myConfiguration.InfluxdbServerUrl,
		AuthToken: myConfiguration.InfluxdbAuthToken,
		Org:       myConfiguration.InfluxdbOrg,
		Bucket:    myConfiguration.InfluxdbBucket,
		BatchSize: 100,
	}
	influxdbCredentials.InitInfluxdb()

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	go influxdb.InfluxWriteFromChannel()
	go SubscribeThread()
	go Process()

	forever := make(chan bool)
	<-forever
}

func Process() {
	for d := range subscribeChannel {
		// Parse JSON message into struct
		var message types.IotMessage
		if err := json.Unmarshal(d.Body, &message); err != nil {
			log.Print(err.Error())
			// Also ack the message if we can't parse it, otherwise we go into an infinite loop of processing corrupt messages
			_ = d.Ack(false)
			return
		}

		// Convert IotMessage to influx point and enqueue for writing
		influxPoint := influxdb2.NewPoint(message.Measurement, message.Tags, message.Fields, message.Time)

		// Influxdb doesn't allow to change the type of a field after the series has been created.
		// Make sure all numbers are floats to prevent type changes during runtime.
		ForceToFloat(influxPoint)

		influxdb.InfluxWritePointsChannel <- influxPoint

		// Ack after queued to be published.
		_ = d.Ack(false)
	}
}

func ForceToFloat(point *write.Point) {
	for _, field := range point.FieldList() {
		floatVal, err := getFloatSwitchOnly(field.Value)
		if err == nil {
			// No error, so the conversion was successful
			field.Value = floatVal
		}
	}
}

// https://stackoverflow.com/questions/20767724/converting-unknown-interface-to-float64-in-golang
func getFloatSwitchOnly(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint:
		return float64(i), nil
	default:
		return math.NaN(), errors.New("Non-numeric type could not be converted to float")
	}
}
