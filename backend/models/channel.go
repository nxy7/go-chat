package models

type Channel struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Capacity    int       `json:"capacity"`
	ActiveUsers []User    `json:"activeUsers"`
	Messages    []Message `json:"messages"`
}

type ChannelInfo interface {
	GetChannelActiveUsers(id string) ([]User, error)
	GetChannelMessageStream(id string) (<-chan Message, error)
}

func GetChannelMessageStream(channelId string) (<-chan Message, error) {
	return nil, nil
}

var MessageChannels = []Channel{{
	Id:          "2s",
	Name:        "20sec",
	Capacity:    10,
	Description: "Messages in this channel last for 20seconds",
	ActiveUsers: make([]User, 0),
	Messages:    make([]Message, 0),
}, {
	Id:          "2m",
	Name:        "2min",
	Capacity:    4,
	Description: "Messages in this channel last for 2min",
	ActiveUsers: make([]User, 0),
	Messages:    make([]Message, 0),
}, {
	Id:          "30m",
	Name:        "30min",
	Capacity:    3,
	Description: "Messages in this channel last for 30min",
	ActiveUsers: make([]User, 0),
	Messages:    make([]Message, 0),
}, {
	Id:          "6h",
	Name:        "6 hours",
	Capacity:    2,
	Description: "Messages in this channel last for 6 hours",
	ActiveUsers: make([]User, 0),
	Messages:    make([]Message, 0),
}, {
	Id:          "1d",
	Name:        "1 day",
	Capacity:    2,
	Description: "Messages in this channel last for 24 hours",
	ActiveUsers: make([]User, 0),
	Messages:    make([]Message, 0),
}, {
	Id:          "1w",
	Name:        "1 week",
	Capacity:    1,
	Description: "Channel that can only be used by one person at a time. Leave messages for other people to see!",
	ActiveUsers: make([]User, 0),
	Messages:    make([]Message, 0),
}}
