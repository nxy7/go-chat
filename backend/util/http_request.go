package util

import (
	"encoding/json"
	"go-chat/models"
	"net/http"
)

func GetLoginData(r *http.Request) (models.LoginData, error) {
	var loginData models.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		return loginData, err
	}
	return loginData, nil
}
