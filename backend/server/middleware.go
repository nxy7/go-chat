package server

import (
	"log"
	"net/http"
)

func MiddlewareChain(h http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

// Guards access to routes that should be protected
func (h *HandlerCtx) LoggedInUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("bearer")
		log.Println("Bearer: ", bearer)
		next.ServeHTTP(w, r)
	}
	// r.Context().Done()
}
