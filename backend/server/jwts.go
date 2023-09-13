package server

import (
	"fmt"
	"net/http"
	"time"

	// _ "crypto/sha256"
	_ "crypto/hmac"

	"github.com/golang-jwt/jwt/v5"
)

// endpoint used to refresh access tokens
func (c *HandlerCtx) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	// first get refresh token from request (should be a part of httpOnly cookie)

	// check if refresh token is valid

	// if it is then generate new access token and return it
	// r.AddCookie()
}

func GenerateRefreshToken(user string, refreshSecret string, signingMethod jwt.SigningMethod) (string, error) {
	rt := jwt.New(signingMethod)
	claims := rt.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(240 * time.Hour)
	claims["user"] = user

	return rt.SignedString([]byte(refreshSecret))
}

func GenerateAccessToken(refreshToken string, accessSecret string, refreshSecret string, signingMethod jwt.SigningMethod) (string, error) {
	rt, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(refreshSecret), nil
	})
	if err != nil {
		return "", err
	}

	expDate, err := rt.Claims.GetExpirationTime()
	if err != nil {
		return "", nil
	}

	if time.Now().Compare(expDate.Time) != -1 {
		return "", fmt.Errorf("Token has expired")
	}

	username, ok := rt.Claims.(jwt.MapClaims)["user"]
	if !ok {
		return "", fmt.Errorf("Token doesn't containe 'user' claim")
	}

	userstring := username.(string)

	at := jwt.New(signingMethod)
	claims := rt.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour)
	claims["user"] = userstring

	return at.SignedString([]byte(accessSecret))
}
