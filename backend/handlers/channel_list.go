package handlers

import (
	"encoding/json"
	"fmt"
	"go-chat/models"
	"log"
	"net/http"
)

// get messages from channel endpoint
func (h *HandlerCtx) ChannelHandler(w http.ResponseWriter, r *http.Request) {
	channels := models.MessageChannels

	for i := range channels {
		listenerCount, err := h.Redis.GetChannelListenersCount(channels[i].Id)
		if err != nil {
			log.Println(err)
			return
		}
		channels[i].ListenerCount = int(listenerCount)
	}

	m, err := json.Marshal(channels)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	w.Write(m)
}
