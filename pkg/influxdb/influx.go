package influxdb

import (
	"context"
	"crypto/tls"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
)

var (
	client           influxdb2.Client
	queryAPI         api.QueryAPI
	writeApiAsync    api.WriteAPI
	writeApiBlocking api.WriteAPIBlocking
	bucket           string

	InfluxWritePointsChannel = make(chan *write.Point)
	InfluxStopWriting        = make(chan struct{})
	InfluxHasStoppedWriting  = make(chan struct{})
)

type InfluxCredentials struct {
	ServerUrl string
	AuthToken string
	Bucket    string
	Org       string
	BatchSize uint
}

func (credentials *InfluxCredentials) InitInfluxdb() {
	if credentials.BatchSize == 0 {
		client = influxdb2.NewClientWithOptions(credentials.ServerUrl, credentials.AuthToken,
			influxdb2.DefaultOptions().
				SetHTTPRequestTimeout(900).
				// default batch size is 5000
				SetTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		)
	} else {
		client = influxdb2.NewClientWithOptions(credentials.ServerUrl, credentials.AuthToken,
			influxdb2.DefaultOptions().
				SetHTTPRequestTimeout(900).
				SetBatchSize(credentials.BatchSize).
				SetTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		)
	}

	queryAPI = client.QueryAPI(credentials.Org)
	writeApiAsync = client.WriteAPI(credentials.Org, credentials.Bucket)
	writeApiBlocking = client.WriteAPIBlocking(credentials.Org, credentials.Bucket)
	bucket = credentials.Bucket

	log.Println("Influxdb2 initialised")
}

func InfluxWriteFromChannel() {
	// Get errors channel
	errorsCh := writeApiAsync.Errors()
	// Create go proc for reading and logging errors
	go func() {
		for err := range errorsCh {
			log.Fatalf("influxdb error: %s\n", err.Error())
		}
	}()

	i := 0
loop:
	for {
		select {
		case <-InfluxStopWriting: // triggered when the stop channel is closed
			log.Println("Stopping influxdb write thread")
			break loop // stop listening to messages
		case point := <-InfluxWritePointsChannel:
			//log.Println("[Influx Write]", point.Time(), utils.PrettyPrint(point.TagList()), utils.PrettyPrint(point.FieldList()))
			i++
			//writeApiAsync.WritePoint(point)
			err := writeApiBlocking.WritePoint(context.Background(), point)
			if err != nil {
				log.Println("failed", err)
				// workaround skip latitudes because this shard had them added as strings
				point.AddField("latitude", string(""))
				err = writeApiBlocking.WritePoint(context.Background(), point)
				log.Println("retry", err.Error())
			}
		}
	}

	log.Printf("Wrote %d points to influxdb", i)

	log.Println("Influxdb flush")
	// Force all unwritten data to be sent
	writeApiAsync.Flush()
	log.Println("Influxdb close")
	// Ensures background processes finishes
	client.Close()

	log.Println("Influxdb write thread stop")
	close(InfluxHasStoppedWriting)
}
