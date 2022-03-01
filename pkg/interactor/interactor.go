package interactor

import (
	"context"

	"github.com/tachunwu/conscious-app/pkg/entity"
)

type Interactor interface {
	// Simple LCRUD
	ListUser(context.Context, entity.ListOptions) ([]entity.User, error)
	CreateUser(context.Context, entity.User) (entity.User, error)
	GetUser(context.Context, entity.Id) (entity.User, error)
	UpdateUser(context.Context, entity.User) (entity.User, error)
	DeleteUser(context.Context, entity.Id) (entity.User, error)

	// More complex example
	CreateTweet(context.Context, entity.Tweet) error
	GetTimeline(context.Context, entity.User) ([]entity.Tweet, error)

	CreateConversation(context.Context, entity.Conversation) (entity.Message, error)
	CreateMessage(context.Context, entity.Message) error
	ListMessage(context.Context, entity.ListOptions) ([]entity.Message, error)
}

type interactor struct {
}
