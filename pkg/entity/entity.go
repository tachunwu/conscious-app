package entity

import "time"

type User struct {
	Id             string
	Name           string
	Email          string
	PasswordDigest string
	CreateAt       time.Time
	UpdateAt       time.Time
}

type Relationship struct {
	Id         string
	FollowerId string
	FollowedId string
	CreateAt   time.Time
	UpdateAt   time.Time
}

type Tweet struct {
	Id       string
	Content  string
	UserId   string
	CreateAt time.Time
	UpdateAt time.Time
}

type Conversation struct {
	Id       string
	UserOne  string
	UserTwo  string
	Status   bool
	CreateAt time.Time
	UpdateAt time.Time
}

type Message struct {
	Id                   string
	IsSeen               bool
	IsDeleteFromSender   bool
	IsDeleteFromReceiver bool
	Content              string
	ConversationId       string
	UserId               string
	CreateAt             time.Time
	UpdateAt             time.Time
}
