package redis

import "go-chat/models"

func (r *Redis) PublishChatMessage(msg models.Message, channelId string) error {
	return nil
}

func (r *Redis) GetChannelMessageStream(channelId string) (<-chan models.Message, error) {
	return nil, nil
}
