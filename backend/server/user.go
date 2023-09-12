package server

import (
	"fmt"
	"net/http"
)

// get user details
func (h *HandlerCtx) UserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// get user with most messages
