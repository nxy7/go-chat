package server

import (
	"go-chat/config"
	"log"
	"net/http"
	"strconv"
)

func Start(c config.Config) error {
	handlerCtx, err := InitializeHandlerCtx(c)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc(c.Prefix+"/health", handlerCtx.HealthHandler)
	mux.HandleFunc(c.Prefix+"/user/{id}/messages", handlerCtx.UserMessagesHandler)
	mux.HandleFunc(c.Prefix+"/user/{id}/details", handlerCtx.UserDetailsHandler)
	mux.HandleFunc(c.Prefix+"/channel", handlerCtx.ChannelHandler)
	mux.HandleFunc(c.Prefix+"/channel/{id}/message", handlerCtx.ChannelMessageHandler)
	mux.HandleFunc(c.Prefix+"/channel/{id}/details", handlerCtx.ChannelDetailsHandler)
	mux.HandleFunc(c.Prefix+"/channel/{id}/eventStream", handlerCtx.ChannelEventStreamHandler)

	log.Println("Starting server on :" + strconv.Itoa(c.Port) + c.Prefix)
	return http.ListenAndServe(":"+strconv.Itoa(c.Port), mux)
}
