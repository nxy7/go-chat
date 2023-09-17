package handlers

import (
	"fmt"
	"net/http"
)

func (h *HandlerCtx) UserMessagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
