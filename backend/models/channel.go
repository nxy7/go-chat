package models

type Channel struct {
	ActiveUsers []User
	Messages    []Message
}

type ChannelInfo interface {
	GetChannelActiveUsers(id string) ([]User, error)
	GetChannelMessageStream(id string) (<-chan Message, error)
}
