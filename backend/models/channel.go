package models

import "time"

type Channel struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	MessageDuration time.Duration
	Description     string `json:"description"`
	Capacity        int    `json:"capacity"`
	ListenerCount   int    `json:"listenerCount"`
}

type ChannelInfo interface {
	GetChannelActiveUsers(id string) ([]User, error)
	GetChannelMessageStream(id string) (<-chan Message, error)
}

var MessageChannels = []Channel{{
	Id:              "2s",
	Name:            "20sec",
	MessageDuration: time.Second * 20,
	Capacity:        10,
	Description:     "Messages in this channel last for 20seconds",
}, {
	Id:              "2m",
	Name:            "2min",
	MessageDuration: time.Minute * 2,
	Capacity:        4,
	Description:     "Messages in this channel last for 2min",
}, {
	Id:              "30m",
	Name:            "30min",
	MessageDuration: time.Minute * 30,
	Capacity:        3,
	Description:     "Messages in this channel last for 30min",
}, {
	Id:              "6h",
	Name:            "6 hours",
	MessageDuration: time.Hour * 6,
	Capacity:        2,
	Description:     "Messages in this channel last for 6 hours",
}, {
	Id:              "1d",
	Name:            "1 day",
	MessageDuration: time.Hour * 24,
	Capacity:        2,
	Description:     "Messages in this channel last for 24 hours",
}, {
	Id:              "1w",
	Name:            "1 week",
	MessageDuration: time.Hour * 24 * 7,
	Capacity:        1,
	Description:     "Channel that can only be used by one person at a time. Leave messages for other people to see!",
}}
