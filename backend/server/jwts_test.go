package server

import (
	"crypto"
	"log"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func Test_JwtTokens(t *testing.T) {
	signingMethod := jwt.SigningMethodHMAC{Hash: crypto.SHA256}
	rt, err := GenerateRefreshToken("noxy", "testsecret", &signingMethod)
	if err != nil {
		t.Error(err)
	} else {
		log.Println(rt)
	}

}
