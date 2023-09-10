package mongodb

import "go-chat/models"

func (db *MongoDb) GetUser(name string) (models.User, error) {
	return models.User{}, nil
}

func (db *MongoDb) AddUser(u models.User) error {
	return nil
}

func (db *MongoDb) AddMessage(u models.User, msg models.Message) error {
	return nil
}

func (db *MongoDb) DeleteUser(name string) error {
	return nil
}
