package server

import (
	"encoding/json"
	"fmt"
	"go-chat/models"
	"net/http"
)

// get messages from channel endpoint
func (h *HandlerCtx) ChannelMessageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// get messages from channel endpoint
func (h *HandlerCtx) ChannelHandler(w http.ResponseWriter, r *http.Request) {
	channels := models.MessageChannels
	m, err := json.Marshal(channels)
	if err != nil {
		fmt.Fprint(w, err)
		// w.Write(err)
		return
	}

	w.Write(m)
}

// get messages from user endpoint
func (h *HandlerCtx) UserMessagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
