package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *HandlerCtx) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, "Wrong method")
		return
	}

	var loginData LoginData
	var bodyBytes []byte
	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if err = json.Unmarshal(bodyBytes, &loginData); err != nil {
		fmt.Fprint(w, err)
		return
	}
	log.Println(loginData)
	fmt.Fprint(w, "log")
}

func (c *HandlerCtx) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, "Wrong method")
		return
	}

	var loginData LoginData
	var bodyBytes []byte
	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if err = json.Unmarshal(bodyBytes, &loginData); err != nil {
		fmt.Fprint(w, err)
		return
	}

}
