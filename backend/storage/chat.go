package storage

import (
	"context"
	"go-chat/models"
)

type MessageStorage interface {
	GetLiveChannelMessages(channelName string, context context.Context) (chan models.Message, error)
	SendMessage(message models.Message, channelName string) error
}
