package storage

import (
	"context"
	"encoding/json"
	"go-chat/models"
	"log"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client redis.Client
}

var _ MessageStorage = (*Redis)(nil)

func (r *Redis) GetLiveChannelMessages(channelName string) (<-chan models.Message, error) {
	p := r.Client.PSubscribe(context.Background(), channelName)
	redisChan := p.Channel()
	messageChannel := make(chan models.Message, 5)
	go func() {
		for v := range redisChan {
			payload := v.Payload
			var asMsg *models.Message
			err := json.Unmarshal([]byte(payload), asMsg)
			if err != nil {
				log.Println(err)
			} else {
				messageChannel <- *asMsg
			}
		}
	}()

	return messageChannel, nil
}

func (r *Redis) SendMessage(message models.Message, channelName string) error {
	err := r.Client.Publish(context.Background(), channelName, message).Err()
	if err != nil {
		return err
	}
	return nil
}
