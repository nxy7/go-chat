// contains message struct and functionality related to sending messages to database and redis
package models

import (
	"time"
)

type Message struct {
	AuthorName string
	Content    string
	Hidden     bool
	CreatedAt  time.Time
}

type MessagePublisher interface {
	PublishChatMessage(msg Message, channelId string) error
	GetChannelMessageStream(id string) (<-chan Message, error)
}
