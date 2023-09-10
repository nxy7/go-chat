package server

import (
	"fmt"
	"net/http"
)

// get user details
func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// get user with most messages
