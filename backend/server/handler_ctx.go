package server

import (
	"context"
	"go-chat/config"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HandlerCtx struct {
	MongoDB mongo.Client
	Redis   redis.Client
}

func InitializeHandlerCtx(c config.Config) (*HandlerCtx, error) {
	uri := "mongodb://" + c.MongoUser + ":" + c.MongoPassword + "@" + c.MongoUrl + "/"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	mc, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	rc := redis.NewClient(&redis.Options{
		Addr: c.RedisUrl,
	})
	return &HandlerCtx{MongoDB: *mc, Redis: *rc}, nil
}
