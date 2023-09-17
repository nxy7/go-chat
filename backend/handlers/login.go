package handlers

import (
	"fmt"
	"go-chat/jwts"
	"go-chat/util"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (c *HandlerCtx) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			fmt.Fprint(w, "Wrong method")
			return
		}

		loginData, err := util.GetLoginData(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("%s", loginData)))
			return
		}
		log.Println(loginData)

		u, err := c.MongoDB.GetUser(loginData.Username)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			w.Write([]byte("There's no user with given login credentials"))
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(loginData.Password))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("There's no user with given login credentials"))
			return
		}

		refreshJwt, err := jwts.GenerateRefreshToken(loginData.Username, c.Config.RefreshTokenSecret, c.Config.JwtSigningMethod)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}
		log.Println(refreshJwt)
		refreshCookie := http.Cookie{
			Name:     "chat-app-refresh",
			Value:    refreshJwt,
			HttpOnly: true,
		}
		http.SetCookie(w, &refreshCookie)

		accessJwt, err := jwts.GenerateAccessToken(refreshJwt, c.Config.AccessTokenSecret, c.Config.RefreshTokenSecret, c.Config.JwtSigningMethod)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint(err)))
			return
		}

		log.Println(loginData, "Registered")
		w.Write([]byte(fmt.Sprint(accessJwt)))
	}
}
