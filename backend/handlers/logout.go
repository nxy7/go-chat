package handlers

import (
	"net/http"
)

func (c *HandlerCtx) LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refreshCookie := http.Cookie{
			Name:     "chat-app-refresh",
			MaxAge:   -1,
			HttpOnly: true,
		}
		http.SetCookie(w, &refreshCookie)
		w.Write([]byte("ok"))
	}
}
