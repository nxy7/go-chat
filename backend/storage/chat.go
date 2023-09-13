package storage

import "go-chat/models"

type MessageStorage interface {
	GetLiveChannelMessages(channelName string) (<-chan models.Message, error)
	SendMessage(message models.Message, channelName string) error
}
