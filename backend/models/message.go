// contains message struct and functionality related to sending messages to database and redis
package models

import (
	"encoding"
	"encoding/json"
)

type Message struct {
	AuthorName string
	Content    string
	Hidden     bool
	CreatedAt  int64
}

func (m Message) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
func (m *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}

var _ encoding.BinaryMarshaler = Message{}
var _ encoding.BinaryUnmarshaler = (*Message)(nil)

type MessagePublisher interface {
	PublishChatMessage(msg Message, channelId string) error
	GetChannelMessageStream(id string) (<-chan Message, error)
}
