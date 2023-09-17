package jwts

import (
	"fmt"
	"log"
	"time"

	// _ "crypto/sha256"
	_ "crypto/hmac"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(user string, refreshSecret string, signingMethod jwt.SigningMethod) (string, error) {
	rt := jwt.New(signingMethod, func(t *jwt.Token) {})
	// rt.Header["typ"] = "JWT"
	// rt.Header["alg"] = signingMethod.Alg()
	claims := rt.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(240 * time.Hour).Unix()
	claims["user"] = user
	log.Println(signingMethod.Alg())
	log.Println(rt)

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
	claims := at.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["user"] = userstring

	return at.SignedString([]byte(accessSecret))
}

func UserIsAuthenticated(accessToken string, accessSecret string) (bool, error) {
	rt, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return false, err
	}
	exp, err := rt.Claims.GetExpirationTime()
	if err != nil {
		return false, err
	}
	if exp.Compare(time.Now()) != 1 {
		return false, fmt.Errorf("Token expired")
	}

	return true, nil
}
