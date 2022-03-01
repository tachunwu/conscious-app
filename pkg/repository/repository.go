package repository

import (
	"context"

	"github.com/tachunwu/conscious-app/pkg/entity"
)

type Repository interface {
	// Simple LCRUD
	ListUser(context.Context, entity.ListOptions) ([]entity.User, error)
	CreateUser(context.Context, entity.User) (entity.User, error)
	GetUser(context.Context, entity.Id) (entity.User, error)
	UpdateUser(context.Context, entity.User) (entity.User, error)
	DeleteUser(context.Context, entity.Id) (entity.User, error)
}

type crdbRepository struct {
}

type postgresRepository struct {
}

type mongoRepository struct {
}
