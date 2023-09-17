package handlers

import (
	"encoding/json"
	"fmt"
	"go-chat/models"
	"log"
	"net/http"
	"strings"
)

func (h *HandlerCtx) ChannelDetailsHandler(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, h.Config.Prefix+"/channel/details/")
	log.Println(slug)

	for _, channel := range models.MessageChannels {
		if channel.Id == slug {
			marchaled, err := json.Marshal(channel)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			} else {
				w.Write(marchaled)
				return
			}
		}
	}

	w.Write([]byte(fmt.Sprint("Channel not found")))

}
