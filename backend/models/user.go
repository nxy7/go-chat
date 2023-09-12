package models

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	Name         string
	Avatar       string
	Messages     []Message
	MessageCount int32
}

type UserStore interface {
	GetUser(name string) (User, error)
	AddMessage(u User, message Message) error
	AddUser(u User) error
}

func GetUser(name string, mdb mongo.Client) (User, error) {
	return User{}, nil
}

func (u *User) UpsertUser(mdb mongo.Client) error {
	return nil
}

func (u *User) UpdateAvatar(avatarUrl string) error {
	return nil
}

// func (u *User) AddMessage(msg Message) error {
// 	return nil
// }

func DeleteUser(name string) error {
	return nil
}
