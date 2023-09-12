package server

import "net/http"

func (h *HandlerCtx) ChannelDetailsHandler(w http.ResponseWriter, r *http.Request) {

}

// Returns stream of events that can include addition, deletion or editing of new messages and user joining or leaving channel
func (h *HandlerCtx) ChannelEventStreamHandler(w http.ResponseWriter, r *http.Request) {
	// events should be retrieved from redis

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
