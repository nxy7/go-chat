package server

import (
	"context"
	"go-chat/config"
	"go-chat/storage"

	// "github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HandlerCtx struct {
	MongoDB storage.Mongo
	// FoundationDB fdb.Database
	Redis  storage.Redis
	Config config.Config
}

func InitializeHandlerCtx(c config.Config) (*HandlerCtx, error) {
	uri := "mongodb://" + c.MongoUser + ":" + c.MongoPassword + "@" + c.MongoUrl + "/"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	mc, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	mDb := mc.Database("chat-app")
	rc := redis.NewClient(&redis.Options{
		Addr: c.RedisUrl,
	})
	return &HandlerCtx{storage.Mongo{Db: mDb}, storage.Redis{Client: *rc}, c}, nil
}
