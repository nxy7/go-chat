package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (h *HandlerCtx) UserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, h.Config.Prefix+"/user/details/")
	// fmt.Printf("username: %v\n", username)
	// token := util.ExtractAccessToken(r)
	// if token == "" {
	// 	w.WriteHeader(401)
	// 	return
	// }
	// at, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(h.Config.AccessTokenSecret), nil
	// })
	// if err != nil {
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }
	// claims := at.Claims.(jwt.MapClaims)
	// username, ok := claims["user"].(string)
	// if !ok {
	// 	return
	// }
	// r.
	u, err := h.MongoDB.GetUser(username)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	toJson, err := json.Marshal(u)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Fprintf(w, string(toJson))
}
