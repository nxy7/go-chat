package mongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDb struct {
	ip       string
	user     string
	password string
}

func (db *MongoDb) GetClient() (*mongo.Client, error) {
	return nil, errors.New("not implemented")
}
