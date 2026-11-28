package service

import (
	"context"

	"nimbus/internal/user"
)

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

type Store interface {
	Register(ctx context.Context, email, passwordHash string) (user.User, error)
}
