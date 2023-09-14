package server

import (
	"encoding/json"
	"fmt"
	"go-chat/models"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func (h *HandlerCtx) ChannelDetailsHandler(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, h.Config.Prefix+"/channel/details/")
	log.Println(slug)

	// channel := models.MessageChannels[slug]
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Returns stream of events that can include addition, deletion or editing of new messages and user joining or leaving channel
func (h *HandlerCtx) ChannelEventStreamHandler(w http.ResponseWriter, r *http.Request) {
	channelId := strings.TrimPrefix(r.URL.Path, "/channel/eventStream/")
	if len(channelId) == 0 {
		w.WriteHeader(402)
		w.Write([]byte("No channel id found"))
		return
	}

	// events should be retrieved from redis
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// Handle error
		w.WriteHeader(402)
		w.Write([]byte(err.Error()))
		return
	}
	defer conn.Close()

	messages, err := h.Redis.GetLiveChannelMessages(channelId)
	if err != nil {
		// Handle error
		w.WriteHeader(402)
		w.Write([]byte(err.Error()))
		return
	}

	for m := range messages {
		conn.WriteJSON(m)
	}

}

func (h *HandlerCtx) PostMessageInChannel(w http.ResponseWriter, r *http.Request) {

}

// While user is keeping this channel open make redis lock
func (h *HandlerCtx) JoinChannel(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in in Redis

	// Lock in user in Redis

	// While user is listening keep redis lock

	// when user leaves notify everyone

}
