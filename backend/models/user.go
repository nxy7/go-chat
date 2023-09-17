package models

type User struct {
	Name         string
	PasswordHash string
	Avatar       string
	Messages     []Message
	MessageCount int32
}

type UserStore interface {
	GetUser(name string) (User, error)
	AddMessage(u User, message Message) error
	AddUser(u User) error
}
