package tts

import (
	"fmt"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"strconv"
	"time"
	"wcrwg-iot-ingress/pkg/constants"
	"wcrwg-iot-ingress/pkg/types"
	"wcrwg-iot-ingress/pkg/utils"
)

func UplinkToGatewayMessage(packetIn ttnpb.ApplicationUp) ([]types.IotMessage, error) {
	var gatewayMessages []types.IotMessage

	for _, gateway := range packetIn.GetUplinkMessage().RxMetadata {
		var gatewayMessage types.IotMessage
		gatewayMessage.Measurement = "gateway"

		AddGatewayMetadataFields(packetIn, gateway, &gatewayMessage)

		gatewayMessages = append(gatewayMessages, gatewayMessage)
	}

	return gatewayMessages, nil
}

func AddGatewayMetadataFields(packetIn ttnpb.ApplicationUp, gatewayIn *ttnpb.RxMetadata, gatewayOut *types.IotMessage) {

	gatewayOut.Time = time.Unix(packetIn.ReceivedAt.Seconds, int64(packetIn.ReceivedAt.Nanos))

	gatewayOut.AddTag(constants.NetworkId,
		constants.NS_TTS_V3+"://"+
			packetIn.GetUplinkMessage().GetNetworkIds().TenantId+"@"+
			utils.NetIdToString(packetIn.GetUplinkMessage().GetNetworkIds().NetId),
	)

	gatewayOut.AddTag(constants.ApplicationId, packetIn.GetEndDeviceIds().ApplicationIds.ApplicationId)
	gatewayOut.AddTag(constants.DeviceId, packetIn.GetEndDeviceIds().DeviceId)

	if packetIn.GetEndDeviceIds().DevEui != nil {
		gatewayOut.AddTag(constants.DeviceEui, fmt.Sprintf("%016X", packetIn.GetEndDeviceIds().DevEui))
	}

	gatewayOut.AddField(constants.FPort, packetIn.GetUplinkMessage().FPort)
	gatewayOut.AddField(constants.FCnt, packetIn.GetUplinkMessage().FCnt)

	gatewayOut.AddField(constants.Frequency, utils.SanitizeFrequency(float64(packetIn.GetUplinkMessage().Settings.Frequency)))

	if packetIn.GetUplinkMessage().Settings.DataRate.GetLora() != nil {
		//log.Println("Is LORA")
		gatewayOut.AddField(constants.Modulation, "LORA")
		gatewayOut.AddField(constants.SpreadingFactor, packetIn.GetUplinkMessage().Settings.DataRate.GetLora().SpreadingFactor)
		gatewayOut.AddField(constants.Bandwidth, packetIn.GetUplinkMessage().Settings.DataRate.GetLora().Bandwidth)
		gatewayOut.AddField(constants.CodingRate, packetIn.GetUplinkMessage().Settings.DataRate.GetLora().CodingRate)
	}
	if packetIn.GetUplinkMessage().Settings.DataRate.GetFsk() != nil {
		//log.Println("Is FSK")
		gatewayOut.AddField(constants.Modulation, "FSK")
		gatewayOut.AddField(constants.Bitrate, packetIn.GetUplinkMessage().Settings.DataRate.GetFsk().BitRate)
	}
	if packetIn.GetUplinkMessage().Settings.DataRate.GetLrfhss() != nil {
		gatewayOut.AddField(constants.Modulation, "LR_FHSS")
		gatewayOut.AddField(constants.Bandwidth, packetIn.GetUplinkMessage().Settings.DataRate.GetLrfhss().GetOperatingChannelWidth())
		gatewayOut.AddField(constants.CodingRate, packetIn.GetUplinkMessage().Settings.DataRate.GetLrfhss().CodingRate)
		// TODO: grid steps, code rate
	}

	// Gateway data

	// Message time, if gateway provides a time, use that rather than the network time
	if gatewayIn.Time != nil {
		gatewayOut.Time = time.Unix(gatewayIn.Time.Seconds, int64(gatewayIn.Time.Nanos))
	}

	// Tags change never per series

	// The gateway's ID - unique per network
	gatewayOut.AddTag(constants.GatewayId, gatewayIn.GatewayIds.GatewayId)
	if gatewayIn.GatewayIds.Eui != nil {
		gatewayOut.AddTag(constants.GatewayEui, fmt.Sprintf("%016X", gatewayIn.GatewayIds.Eui))
	}

	if gatewayIn.PacketBroker != nil {
		/*
			ttsDomain = forwrder_tenant_id@forwarder_net_id // "ttn@000013"
		*/
		forwarderTenantId := gatewayIn.PacketBroker.ForwarderTenantId
		forwarderNetId := gatewayIn.PacketBroker.ForwarderNetId
		if forwarderTenantId == "ttnv2" {
			gatewayOut.AddTag(constants.NetworkId, "thethingsnetwork.org")
		} else {
			gatewayOut.AddTag(constants.NetworkId, constants.NS_TTS_V3+"://"+forwarderTenantId+"@"+utils.NetIdToString(forwarderNetId))
			gatewayOut.AddTag(constants.ClusterId, gatewayIn.PacketBroker.ForwarderClusterId)
		}

		/*
			Use GatewayId and EUI if reported by PacketBroker
		*/
		if gatewayIn.PacketBroker.ForwarderGatewayEui != nil {
			gatewayOut.AddTag(constants.GatewayEui, fmt.Sprintf("%016X", gatewayIn.PacketBroker.ForwarderGatewayEui))
		}
		if gatewayIn.PacketBroker.ForwarderGatewayId != nil {
			gatewayOut.AddTag(constants.GatewayId, gatewayIn.PacketBroker.ForwarderGatewayId.Value)
		}

	} else {
		gatewayOut.AddTag(constants.NetworkId, constants.NS_TTS_V3+"://"+
			packetIn.GetUplinkMessage().NetworkIds.TenantId+"@"+
			utils.NetIdToString(packetIn.GetUplinkMessage().NetworkIds.NetId))
		gatewayOut.AddTag(constants.ClusterId, packetIn.GetUplinkMessage().NetworkIds.ClusterId)
	}

	gatewayOut.AddTag(constants.AntennaIndex, strconv.Itoa(int(gatewayIn.AntennaIndex)))

	// Fields - changes often (in every uplink)
	gatewayOut.AddField(constants.ChannelIndex, gatewayIn.ChannelIndex)

	if gatewayIn.ChannelRssi != 0 {
		gatewayOut.AddField(constants.Rssi, gatewayIn.ChannelRssi)
	}
	if gatewayIn.Rssi != 0 {
		gatewayOut.AddField(constants.Rssi, gatewayIn.Rssi)
	}
	if gatewayIn.SignalRssi != nil {
		gatewayOut.AddField(constants.SignalRssi, gatewayIn.SignalRssi.Value)
	}
	gatewayOut.AddField(constants.Snr, gatewayIn.Snr)
	if gatewayIn.Location != nil {
		gatewayOut.AddField(constants.Latitude, gatewayIn.Location.Latitude)
		gatewayOut.AddField(constants.Longitude, gatewayIn.Location.Longitude)
		gatewayOut.AddField(constants.Altitude, gatewayIn.Location.Altitude)
		gatewayOut.AddField(constants.LocationAccuracy, gatewayIn.Location.Accuracy)

		gatewayOut.AddField(constants.LocationSource, gatewayIn.Location.Source.String())
	}
}
