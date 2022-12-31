package tts

import (
	"github.com/go-chi/render"
	"github.com/golang/protobuf/proto"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

/*
TTS V3 Webhook
Header TTNMAPPERORG-USER contains the email address of the user for identification of the source of the data
Header TTNMAPPERORG-EXPERIMENT indicates if mapping is done to an experiment, and the experiment name
*/
func (handlerContext *Context) PostV3Uplink(w http.ResponseWriter, r *http.Request) {
	i := strconv.Itoa(rand.Intn(100))

	// Read data
	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print("[" + i + "] " + err.Error())
		return
	}
	var packetIn ttnpb.ApplicationUp

	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		marshaler := jsonpb.TTN()
		if err := marshaler.Unmarshal(body, &packetIn); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response["success"] = false
			response["message"] = "Can not parse json body"
			log.Print("[" + i + "] " + err.Error())
			return
		}
	} else if contentType == "application/protobuf" || contentType == "application/x-protobuf" || contentType == "application/octet-stream" { // TTS uses application/octet-stream
		if err := proto.Unmarshal(body, &packetIn); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response["success"] = false
			response["message"] = "Can not parse protobuf body"
			log.Print("[" + i + "] " + err.Error())
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Content-Type header not set"
		return
	}

	if packetIn.GetUplinkMessage() == nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "uplink_message not set"
		log.Print("[" + i + "] uplink_message not set")
		return
	}

	if packetIn.GetUplinkMessage().FPort == 0 {
		response["success"] = false
		response["message"] = "fPort is 0"
		log.Print("[" + i + "] fPort is 0")
		return
	}

	if packetIn.GetUplinkMessage().DecodedPayload == nil {
		response["success"] = false
		response["message"] = "payload_fields not set"
		log.Print("[" + i + "] payload_fields not set")
		//return // Do not return, as we can still use the metadata to update gateway last seen and contribute channels and signal stats
	}

	if packetIn.GetUplinkMessage().GetNetworkIds() == nil {
		response["success"] = false
		response["message"] = "Network IDs not set"
		log.Print("[" + i + "] Network IDs not set")
		return
	}

	// 1. Sensor data
	sensorMessage, err := UplinkToSensorMessage(packetIn)
	if err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print("[" + i + "] " + err.Error())
		return
	}

	// Push this new packet into the stack
	handlerContext.PublishChannel <- sensorMessage

	// 2. Gateway data
	gatewayMessages, err := UplinkToGatewayMessage(packetIn)
	if err != nil {
		response["success"] = false
		response["message"] = err.Error()
		log.Print("[" + i + "] " + err.Error())
		return
	}

	for _, gatewayMessage := range gatewayMessages {
		handlerContext.PublishChannel <- gatewayMessage
	}

	w.WriteHeader(http.StatusAccepted)
	response["success"] = true
	response["message"] = "New packet accepted into queue"
}

func (handlerContext *Context) GetV3(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["success"] = true
	response["message"] = "GET test success"
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, response)
}

func (handlerContext *Context) PostV3JoinAccept(w http.ResponseWriter, r *http.Request) {

	// Read data
	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print(err.Error())
		return
	}

	log.Println("Join Accept: ", string(body))

	// TODO implement this endpoint
	w.WriteHeader(http.StatusOK)
}

func (handlerContext *Context) PostV3LocationSolved(w http.ResponseWriter, r *http.Request) {

	// Read data
	response := make(map[string]interface{})
	defer render.JSON(w, r, response)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response["success"] = false
		response["message"] = "Can not read POST body"
		log.Print(err.Error())
		return
	}

	log.Println("Location Solved: ", string(body))

	// TODO implement this endpoint
	w.WriteHeader(http.StatusOK)

}
