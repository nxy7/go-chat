package config

import (
	_ "crypto/sha256"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Config struct {
	Prefix        string
	Port          int
	RedisUrl      string
	MongoUrl      string
	MongoUser     string
	MongoPassword string

	RefreshTokenSecret string
	AccessTokenSecret  string
	JwtSigningMethod   *jwt.SigningMethodHMAC
}

func FromEnv() Config {
	// .env file is used only locally
	_ = godotenv.Load("../.env")

	prefix := os.Getenv("BACKEND_PREFIX")
	port := 8282
	redisUrl := os.Getenv("REDIS_URL")
	mongoUrl := os.Getenv("MONGO_URL")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	signingMethod := jwt.SigningMethodHS256

	return Config{
		Prefix:             prefix,
		Port:               port,
		RedisUrl:           redisUrl,
		MongoUrl:           mongoUrl,
		MongoUser:          mongoUser,
		MongoPassword:      mongoPassword,
		RefreshTokenSecret: refreshSecret,
		AccessTokenSecret:  accessSecret,
		JwtSigningMethod:   signingMethod,
	}
}
