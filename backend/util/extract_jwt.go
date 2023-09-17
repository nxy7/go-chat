package util

import (
	"net/http"
)

func ExtractAccessToken(r *http.Request) string {
	return r.Header.Get("bearer")
}
