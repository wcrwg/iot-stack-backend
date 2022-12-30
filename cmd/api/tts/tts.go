package tts

import (
	"github.com/go-chi/chi"
	_ "time"
	"wcrwg-iot-ingress/pkg/types"
)

type Context struct {
	PublishChannel chan types.IotMessage
}

func Routes(publishChannel chan types.IotMessage) *chi.Mux {
	context := &Context{PublishChannel: publishChannel}

	router := chi.NewRouter()

	router.Post("/v3/uplink-message", context.PostV3Uplink)
	router.Post("/v3/join-accept", context.PostV3JoinAccept)
	router.Post("/v3/location-solved", context.PostV3LocationSolved)
	router.Get("/v3", context.GetV3)

	return router
}
