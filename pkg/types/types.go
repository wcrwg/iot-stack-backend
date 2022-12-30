package types

import "time"

type IotMessage struct {
	Time        time.Time
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
}
