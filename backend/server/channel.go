package server

import "net/http"

func GetChannelDetails(w http.ResponseWriter, r *http.Request) {

}

// Returns stream of events that can include addition, deletion or editing of new messages and user joining or leaving channel
func GetChannelEventStream(w http.ResponseWriter, r *http.Request) {
	// events should be retrieved from redis

}

func PostMessageInChannel(w http.ResponseWriter, r *http.Request) {

}

// While user is keeping this channel open make redis lock
func JoinChannel(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in in Redis

	// Lock in user in Redis

	// While user is listening keep redis lock

	// when user leaves notify everyone

}
