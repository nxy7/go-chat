package server

import (
	"fmt"
	"go-chat/config"
	"net/http"
)

type Server struct {
	Prefix string
	Port   int32
}

func (s *Server) Start() error {
	http.HandleFunc(s.Prefix+"/health", HealthHandler)
	http.HandleFunc(s.Prefix+"/user/:id/messages", GetUserMessagesHandler)
	http.HandleFunc(s.Prefix+"/user/:id/details", GetUserDetails)
	http.HandleFunc(s.Prefix+"/channel/:id/messages", GetChannelMessagesHandler)
	http.HandleFunc(s.Prefix+"/channel/:id/details", GetChannelDetails)

	http.HandleFunc(s.Prefix+"/channel/:id/eventStream", GetChannelEventStream)

	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
	return nil
}

func FromConfig(config config.Config) Server {
	return Server{}
}
