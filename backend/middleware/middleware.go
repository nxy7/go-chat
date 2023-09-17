package middleware

import (
	"go-chat/handlers"
	"log"
	"net/http"
)

func MiddlewareChain(h http.HandlerFunc, hCtx *handlers.HandlerCtx, middlewares ...func(http.HandlerFunc) http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

// Guards access to routes that should be protected
func LoggedInUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("bearer")
		log.Println("Bearer: ", bearer)
		next.ServeHTTP(w, r)
	}
}
