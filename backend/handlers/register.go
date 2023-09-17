package handlers

import (
	"fmt"
	"go-chat/jwts"
	"go-chat/models"
	"go-chat/util"
	"log"
	"net/http"
)

func (c *HandlerCtx) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Wrong method")
		return
	}

	loginData, err := util.GetLoginData(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	if len(loginData.Username) < 4 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint("Username must be at least 4 characters long")))
		return
	}
	if len(loginData.Password) < 6 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint("Password must be at least 6 characters long")))
		return
	}

	u, err := c.MongoDB.GetUser(loginData.Username)
	if u != nil {
		log.Println(u)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint("User already exists")))
		return
	}

	hashedPassword, err := util.HashPassword(loginData.Password)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	u = &models.User{
		Name:         loginData.Username,
		PasswordHash: hashedPassword,
		Avatar:       "",
		Messages:     []models.Message{},
		MessageCount: 0,
	}
	if err = c.MongoDB.UpsertUser(u); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	refreshJwt, err := jwts.GenerateRefreshToken(loginData.Username, c.Config.RefreshTokenSecret, c.Config.JwtSigningMethod)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	refreshCookie := http.Cookie{
		Name:     "chat-app-refresh",
		Value:    refreshJwt,
		HttpOnly: true,
	}
	http.SetCookie(w, &refreshCookie)

	accessJwt, err := jwts.GenerateAccessToken(refreshJwt, c.Config.AccessTokenSecret, c.Config.RefreshTokenSecret, c.Config.JwtSigningMethod)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	log.Println(loginData, "Registered")
	w.Write([]byte(fmt.Sprint(accessJwt)))
}
