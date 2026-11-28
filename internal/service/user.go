package service

import (
	"context"
	"errors"
	"fmt"

	"nimbus/internal/auth"
	"nimbus/internal/storage"
	"nimbus/internal/user"
)

func (s *Service) Register(ctx context.Context, email, password string) (user.User, error) {
	passwordHash, err := auth.Hash(password)
	if err != nil {
		return user.User{}, fmt.Errorf("hash password: %w", err)
	}

	u, err := s.store.Register(ctx, email, passwordHash)

	if errors.Is(err, storage.ErrEmailAlreadyExists) {
		return user.User{}, ErrEmailAlreadyExists
	}

	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
