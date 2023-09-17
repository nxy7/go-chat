package handlers

import (
	"fmt"
	"go-chat/jwts"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (c *HandlerCtx) RefreshTokenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			fmt.Fprint(w, "Wrong method")
			return
		}

		cookies := r.Cookies()
		var refreshCookie *http.Cookie
		for _, c := range cookies {
			if c.Name == "chat-app-refresh" {
				refreshCookie = c
				break
			}
		}
		refreshJwt := refreshCookie.Value
		parsedJwt, err := jwt.Parse(refreshJwt, func(t *jwt.Token) (interface{}, error) {
			return []byte(c.Config.RefreshTokenSecret), nil
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		exp, err := parsedJwt.Claims.GetExpirationTime()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		if exp.Compare(time.Now()) != 1 {
			w.Write([]byte("Refresh token expired"))
			return
		}

		at, err := jwts.GenerateAccessToken(refreshJwt, c.Config.AccessTokenSecret, c.Config.RefreshTokenSecret, c.Config.JwtSigningMethod)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(at))
	}
}
