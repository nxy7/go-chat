package util

import (
	"encoding/json"
	"go-chat/models"
	"net/http"
)

func GetLoginData(r *http.Request) (models.LoginData, error) {
	var loginData models.LoginData
	var bodyBytes []byte
	_, err := r.Body.Read(bodyBytes)
	if err != nil {
		return loginData, err
	}
	if err = json.Unmarshal(bodyBytes, &loginData); err != nil {
		return loginData, err
	}
	return loginData, nil
}
