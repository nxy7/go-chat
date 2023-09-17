package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"go-chat/models"
	"log"
	"sort"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client redis.Client
}

var _ MessageStorage = (*Redis)(nil)

func (r *Redis) GetChannelOldMessages(channelName string) ([]models.Message, error) {
	keys, err := r.Client.Keys(context.Background(), channelName+":*").Result()
	if err != nil {
		return nil, err
	}
	messages := make([]models.Message, 0)
	for _, k := range keys {
		var m models.Message
		v, err := r.Client.Get(context.Background(), k).Result()
		if err != nil {
			log.Println(err)
			continue
		}
		err = json.Unmarshal([]byte(v), &m)
		if err != nil {
			log.Println(err)
			continue
		}
		messages = append(messages, m)
	}
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt < messages[j].CreatedAt
	})
	return messages, nil
}

func (r *Redis) GetChannelListenersCount(channelId string) (int64, error) {
	counts, err := r.Client.PubSubNumSub(context.Background(), channelId).Result()
	if err != nil {
		return 0, err
	}
	channelCount, ok := counts[channelId]
	if !ok {
		return 0, fmt.Errorf("Function might have received wrong channel name: %s", channelId)
	}

	return channelCount, nil
}

func (r *Redis) GetLiveChannelMessages(channelName string, context context.Context) (chan models.Message, error) {
	p := r.Client.Subscribe(context, channelName)
	redisChan := p.Channel()
	messageChannel := make(chan models.Message, 5)
	go func() {
		<-context.Done()
		log.Println("Closing redis pubsub connection")
		p.Close()
	}()
	go func() {
		for v := range redisChan {
			payload := v.Payload
			var asMsg models.Message
			err := json.Unmarshal([]byte(payload), &asMsg)
			if err != nil {
				log.Println(err)
			} else {
				messageChannel <- asMsg
			}
		}

	}()

	return messageChannel, nil
}

func (r *Redis) SendMessage(message models.Message, channelName string) error {
	var selectedChannel *models.Channel
	for _, c := range models.MessageChannels {
		if c.Id == channelName {
			selectedChannel = &c
			break
		}
	}
	if selectedChannel == nil {
		return fmt.Errorf("Message channel doesn't exist. Channel name: %s", channelName)
	}
	r.Client.Set(context.Background(), channelName+":"+message.AuthorName+":"+strconv.Itoa(int(message.CreatedAt)), message, selectedChannel.MessageDuration)

	err := r.Client.Publish(context.Background(), channelName, message).Err()
	if err != nil {
		return err
	}
	return nil
}
