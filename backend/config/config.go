package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Prefix        string
	Port          int
	RedisUrl      string
	MongoUrl      string
	MongoUser     string
	MongoPassword string
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

	return Config{
		Prefix:        prefix,
		Port:          port,
		RedisUrl:      redisUrl,
		MongoUrl:      mongoUrl,
		MongoUser:     mongoUser,
		MongoPassword: mongoPassword,
	}
}
