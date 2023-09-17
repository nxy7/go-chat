package server

import (
	"go-chat/config"
	"go-chat/handlers"
	"log"
	"net/http"
	"strconv"
)

func Start(c config.Config) error {
	handlerCtx, err := handlers.InitializeHandlerCtx(c)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc(c.Prefix+"/health", handlerCtx.HealthHandler)
	mux.HandleFunc(c.Prefix+"/login", handlerCtx.LoginHandler())
	mux.HandleFunc(c.Prefix+"/register", handlerCtx.RegisterHandler)
	mux.HandleFunc(c.Prefix+"/token", handlerCtx.RefreshTokenHandler())
	mux.HandleFunc(c.Prefix+"/user/details/", handlerCtx.UserDetailsHandler)
	mux.HandleFunc(c.Prefix+"/channel", handlerCtx.ChannelHandler)
	mux.HandleFunc(c.Prefix+"/channel/details/", handlerCtx.ChannelDetailsHandler)
	mux.HandleFunc(c.Prefix+"/channel/eventStream/", handlerCtx.ChannelEventStreamHandler)

	log.Println("Starting server on :" + strconv.Itoa(c.Port) + c.Prefix)
	return http.ListenAndServe(":"+strconv.Itoa(c.Port), mux)
}
