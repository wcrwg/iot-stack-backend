package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

func NetIdToString(netId []byte) string {
	return fmt.Sprintf("%06X", netId)
}

// Some single channel gateways send frequency in Hz, not MHz
// TTNv3 also sends the frequency in Herz, not MHz like V2 - change to Hz here
// Below 1MHz assume the value is passed in MHz not Hz, so convert to Hz
func SanitizeFrequency(frequency float64) uint64 {
	// 868.1 to 868100000 - but we will lose the decimals
	if frequency < 1000.0 {
		frequency = frequency * 1000000
	}

	// 868400000000000 to 868400000
	if frequency > 1000000000 {
		frequency = frequency / 1000000
	}

	// 869099976 to 869100000
	frequency = math.Round(frequency/1000) * 1000
	frequencyInt := uint64(frequency)

	return frequencyInt
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
