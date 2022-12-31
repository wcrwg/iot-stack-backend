package tts

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"log"
	"time"
	"wcrwg-iot-ingress/pkg/constants"
	"wcrwg-iot-ingress/pkg/types"
	"wcrwg-iot-ingress/pkg/utils"
)

func UplinkToSensorMessage(packetIn ttnpb.ApplicationUp) (types.IotMessage, error) {
	var sensorMessage types.IotMessage
	sensorMessage.Measurement = "sensor"

	// Add metadata fields
	AddNetworkMetadataFields(packetIn, &sensorMessage)

	// Store the raw payload too
	sensorMessage.AddField("frm_payload", b64.StdEncoding.EncodeToString(packetIn.GetUplinkMessage().FrmPayload))

	// Add payload fields, and
	// 3. If payload fields are available, try getting coordinates from there
	if packetIn.GetUplinkMessage().DecodedPayload != nil {
		if err := AddParsedPayloadFields(packetIn, &sensorMessage); err != nil {
			return sensorMessage, err
		}
	}

	return sensorMessage, nil
}

func AddNetworkMetadataFields(packetIn ttnpb.ApplicationUp, packetOut *types.IotMessage) {

	/*
		V3
		  "received_at" : "2020-02-12T15:15..."      // ISO 8601 UTC timestamp at which the message has been received by the Application Server
		  "uplink_message" : {
		    "rx_metadata": [{                        // A list of metadata for each antenna of each gateway that received this message
		      "time": "2020-02-12T15:15:45.787Z",    // ISO 8601 UTC timestamp at which the uplink has been received by the gateway
		    }],
		    "settings": {                            // Settings for the transmission
		      "time": "2020-02-12T15:15:45.787Z"     // ISO 8601 UTC timestamp at which the uplink has been received by the gateway
		    },
		    "received_at": "2020-02-12T15:15..."     // ISO 8601 UTC timestamp at which the uplink has been received by the Network Server
	*/
	packetOut.Time = time.Unix(packetIn.ReceivedAt.Seconds, int64(packetIn.ReceivedAt.Nanos))

	/*
		V3
		  "end_device_ids" : {
		    "device_id" : "dev1",                    // Device ID
		    "application_ids" : {
		      "application_id" : "app1"              // Application ID
		    },
		    "dev_eui" : "0004A30B001C0530",          // DevEUI of the end device
		    "join_eui" : "800000000000000C",         // JoinEUI of the end device (also known as AppEUI in LoRaWAN versions below 1.1)
		    "dev_addr" : "00BCB929"                  // Device address known by the Network Server
		  },
	*/

	packetOut.AddTag(constants.NetworkId,
		constants.NS_TTS_V3+"://"+
			packetIn.GetUplinkMessage().GetNetworkIds().TenantId+"@"+
			utils.NetIdToString(packetIn.GetUplinkMessage().GetNetworkIds().NetId),
	)

	log.Printf("[TTS] %s - %s", packetIn.GetEndDeviceIds().ApplicationIds.ApplicationId, packetIn.GetEndDeviceIds().DeviceId)
	packetOut.AddTag(constants.ApplicationId, packetIn.GetEndDeviceIds().ApplicationIds.ApplicationId)
	packetOut.AddTag(constants.DeviceId, packetIn.GetEndDeviceIds().DeviceId)

	if packetIn.GetEndDeviceIds().DevEui != nil {
		packetOut.AddTag(constants.DeviceEui, fmt.Sprintf("%016X", packetIn.GetEndDeviceIds().DevEui))
	}

	/*
	   V3
	         "uplink_message":{
	            "f_port":1,
	            "f_cnt":527,
	*/
	packetOut.AddField(constants.FPort, packetIn.GetUplinkMessage().FPort)
	packetOut.AddField(constants.FCnt, packetIn.GetUplinkMessage().FCnt)

	/*
	   V3
	      "uplink_message" : {
	        "settings": {                            // Settings for the transmission
	          "data_rate": {                         // Data rate settings
	            "lora": {                            // LoRa modulation settings
	              "bandwidth": 125000,               // Bandwidth (Hz)
	              "spreading_factor": 7              // Spreading factor
	            }
	          },
	          "data_rate_index": 5,                  // LoRaWAN data rate index
	          "coding_rate": "4/6",                  // LoRa coding rate
	          "frequency": "868300000",              // Frequency (Hz)
	        },
	*/
	packetOut.AddField(constants.Frequency, utils.SanitizeFrequency(float64(packetIn.GetUplinkMessage().Settings.Frequency)))

	if packetIn.GetUplinkMessage().Settings.DataRate.GetLora() != nil {
		//log.Println("Is LORA")
		packetOut.AddField(constants.Modulation, "LORA")
		packetOut.AddField(constants.SpreadingFactor, packetIn.GetUplinkMessage().Settings.DataRate.GetLora().SpreadingFactor)
		packetOut.AddField(constants.Bandwidth, packetIn.GetUplinkMessage().Settings.DataRate.GetLora().Bandwidth)
		packetOut.AddField(constants.CodingRate, packetIn.GetUplinkMessage().Settings.DataRate.GetLora().CodingRate)
	}
	if packetIn.GetUplinkMessage().Settings.DataRate.GetFsk() != nil {
		//log.Println("Is FSK")
		packetOut.AddField(constants.Modulation, "FSK")
		packetOut.AddField(constants.Bitrate, packetIn.GetUplinkMessage().Settings.DataRate.GetFsk().BitRate)
	}
	if packetIn.GetUplinkMessage().Settings.DataRate.GetLrfhss() != nil {
		packetOut.AddField(constants.Modulation, "LR_FHSS")
		packetOut.AddField(constants.Bandwidth, packetIn.GetUplinkMessage().Settings.DataRate.GetLrfhss().GetOperatingChannelWidth())
		packetOut.AddField(constants.CodingRate, packetIn.GetUplinkMessage().Settings.DataRate.GetLrfhss().CodingRate)
		// TODO: grid steps, code rate
	}

	// 1. Try to use the location from the metadata first. This is likely the location set on the console.
	if packetIn.GetUplinkMessage().Locations["user"] != nil {
		packetOut.AddField(constants.Latitude, packetIn.GetUplinkMessage().Locations["user"].Latitude)
		packetOut.AddField(constants.Longitude, packetIn.GetUplinkMessage().Locations["user"].Longitude)
		packetOut.AddField(constants.Altitude, packetIn.GetUplinkMessage().Locations["user"].Altitude)
		packetOut.AddField(constants.LocationAccuracy, packetIn.GetUplinkMessage().Locations["user"].Accuracy)
		packetOut.AddField(constants.LocationSource, packetIn.GetUplinkMessage().Locations["user"].Source.String())
	}

	// 2. If the packetIn contains a solved location, rather use that - this is sent to the /location-solved endpoint, so useless here
	if packetIn.GetLocationSolved() != nil {
		packetOut.AddField(constants.Latitude, packetIn.GetLocationSolved().Location.Latitude)
		packetOut.AddField(constants.Longitude, packetIn.GetLocationSolved().Location.Longitude)
		packetOut.AddField(constants.Altitude, packetIn.GetLocationSolved().Location.Altitude)
		packetOut.AddField(constants.LocationAccuracy, packetIn.GetLocationSolved().Location.Accuracy)
		packetOut.AddField(constants.LocationSource, packetIn.GetLocationSolved().Location.Source.String())
	}
}

func AddParsedPayloadFields(packetIn ttnpb.ApplicationUp, packetOut *types.IotMessage) error {
	// DecodedPayload is &Struct{Fields:map[string]*Value{},XXX_unrecognized:[],}.
	// Convert to a more standard map[string]interface{}

	// Marshal struct to json
	marshaler := jsonpb.TTN()
	decodedJson, err := marshaler.Marshal(packetIn.GetUplinkMessage().DecodedPayload)
	if err != nil {
		return err
	}

	// Unmarshal json to interface{}
	var decodedInterface map[string]interface{}
	err = json.Unmarshal(decodedJson, &decodedInterface)
	if err != nil {
		return err
	}

	// For each field in the decoded payload, add to IotMessage
	for key, val := range decodedInterface {
		packetOut.AddField(key, val)
	}

	return nil
}
