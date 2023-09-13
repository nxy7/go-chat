package util

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateSalt() string {
	return RandomString(15)
}

func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", fmt.Errorf("At least one input was empty. PasswordLength: %d ", len(password))
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[seededRand.Intn(len(letterRunes))]
	}
	return string(b)
}
