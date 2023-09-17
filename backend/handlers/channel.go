package handlers

import (
	"go-chat/models"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type WebsocketEvent struct {
	Type    string
	Payload any
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *HandlerCtx) ChannelEventStreamHandler(w http.ResponseWriter, r *http.Request) {
	channelId := strings.TrimPrefix(r.URL.Path, h.Config.Prefix+"/channel/eventStream/")
	if len(channelId) == 0 {
		w.WriteHeader(402)
		w.Write([]byte("No channel id found"))
		return
	}
	ctx := r.Context()

	var mChan *models.Channel
	for _, c := range models.MessageChannels {
		if c.Id == channelId {
			mChan = &c
			break
		}
	}
	if mChan == nil {
		w.WriteHeader(405)
		log.Println([]byte("No channel found"))
		w.Write([]byte("No channel found"))
		return
	}
	i, err := h.Redis.GetChannelListenersCount(mChan.Id)
	if err != nil {
		w.WriteHeader(405)
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	if i >= int64(mChan.Capacity) {
		w.WriteHeader(405)
		log.Println("Channel full")
		w.Write([]byte("Channel full"))
		return
	}

	messages, err := h.Redis.GetLiveChannelMessages(channelId, ctx)
	if err != nil {
		w.WriteHeader(402)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	defer conn.Close()

	oldMessages, err := h.Redis.GetChannelOldMessages(channelId)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for m := range messages {
			select {
			case <-ctx.Done():
				log.Println("context done")
				return
			default:
				err = conn.WriteJSON(m)
				if err != nil {
					return
				}
			}
		}
	}()

	for _, m := range oldMessages {
		err = conn.WriteJSON(m)
		if err != nil {
			return
		}
	}
	for {
		fromClientMessage := models.Message{}
		err := conn.ReadJSON(&fromClientMessage)
		if err != nil {
			log.Println(err)
			return
		}
		if fromClientMessage.Content == "" {
			continue
		}
		log.Println(fromClientMessage)
		err = h.Redis.SendMessage(fromClientMessage, channelId)
		if err != nil {
			log.Println(err)
		}
		err = h.MongoDB.UserIncrementMessageCount(fromClientMessage.AuthorName)
		if err != nil {
			log.Println(err)
		}
	}
}
