package storage

import (
	"context"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Db *mongo.Database
}

var _ UserStorage = (*Mongo)(nil)

const UserCollectionName = "users"

func (m *Mongo) GetUser(user string) (*models.User, error) {
	var u models.User
	err := m.Db.Collection(UserCollectionName).FindOne(context.Background(), bson.D{{Key: "name", Value: user}}).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (m *Mongo) UpsertUser(user *models.User) error {
	upsert := true
	_, err := m.Db.Collection(UserCollectionName).UpdateOne(context.Background(), bson.D{{Key: "name", Value: user.Name}}, bson.D{{Key: "$set", Value: user}}, &options.UpdateOptions{
		Upsert: &upsert,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) DeleteUser(user string) error {
	_, err := m.Db.Collection(UserCollectionName).DeleteOne(context.Background(), bson.D{{Key: "name", Value: user}})
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) UserIncrementMessageCount(user string) error {
	var u models.User
	err := m.Db.Collection(UserCollectionName).FindOne(context.Background(), bson.D{{Key: "name", Value: user}}).Decode(&u)
	if err != nil {
		return err
	}
	u.MessageCount++
	_, err = m.Db.Collection(UserCollectionName).UpdateOne(context.Background(), bson.D{{"name", user}}, bson.D{{"$set", u}})
	if err != nil {
		return err
	}

	return nil
}
