package server

import (
	"fmt"
	"net/http"
)

// get messages from channel endpoint
func GetChannelMessagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// get messages from user endpoint
func GetUserMessagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
