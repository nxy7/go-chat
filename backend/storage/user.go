package storage

import "go-chat/models"

type UserStorage interface {
	GetUser(user string) (models.User, error)
	UpsertUser(user models.User) error
	DeleteUser(user string) error
}
