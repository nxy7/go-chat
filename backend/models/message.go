// contains message struct and functionality related to sending messages to database and redis
package models

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Message struct {
	MessageId  string
	AuthorName string
	Content    string
	Hidden     bool
	CreatedAt  time.Time
}

type MessagePublisher interface {
	PublishChatMessage(msg Message, channelId string) error
	GetChannelMessageStream(id string) (<-chan Message, error)
}

func (m *Message) PublishChatMessage(channelId string, redis redis.Client) error {
	// publish message to redis
	// update message count in mongodb
	return nil
}
