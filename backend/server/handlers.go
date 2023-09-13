package server

import (
	"fmt"
	"go-chat/models"
	"go-chat/util"
	"log"
	"net/http"
)

// logins user and redirects to given redirect location
func (c *HandlerCtx) LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	hashedPassword, err := util.HashPassword(loginData.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	// get hashed password from db
	u, err := c.MongoDB.GetUser(loginData.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	// compare password with hashed one
	if hashedPassword != u.PasswordHash {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint("Incorrect password")))
		return
	}

	refreshJwt, err := GenerateRefreshToken(loginData.Username, c.Config.RefreshTokenSecret, &c.Config.JwtSigningMethod)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	refreshCookie := http.Cookie{
		Name:     "chat-app-refresh",
		Value:    refreshJwt,
		HttpOnly: true,
	}
	r.AddCookie(&refreshCookie)

	accessJwt, err := GenerateAccessToken(refreshJwt, c.Config.AccessTokenSecret, c.Config.RefreshTokenSecret, &c.Config.JwtSigningMethod)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	log.Println(loginData, "Registered")
	w.Write([]byte(fmt.Sprint(accessJwt)))
}

// registers user and redirects to given redirect location
func (c *HandlerCtx) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Wrong method")
		return
	}

	loginData, err := util.GetLoginData(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf("%s", loginData)))
		return
	}
	// check if user is already registered
	_, err = c.MongoDB.GetUser(loginData.Username)
	if err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint("User already exists")))
		return
	}

	hashedPassword, err := util.HashPassword(loginData.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	// if not add him to the database
	u := models.User{
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

	refreshJwt, err := GenerateRefreshToken(loginData.Username, c.Config.RefreshTokenSecret, &c.Config.JwtSigningMethod)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	refreshCookie := http.Cookie{
		Name:     "chat-app-refresh",
		Value:    refreshJwt,
		HttpOnly: true,
	}
	r.AddCookie(&refreshCookie)

	accessJwt, err := GenerateAccessToken(refreshJwt, c.Config.AccessTokenSecret, c.Config.RefreshTokenSecret, &c.Config.JwtSigningMethod)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprint(err)))
		return
	}

	log.Println(loginData, "Registered")
	w.Write([]byte(fmt.Sprint(accessJwt)))
}
