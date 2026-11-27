package postgres

import (
	"context"
	"errors"

	"nimbus/internal/storage"
	"nimbus/internal/user"

	"github.com/jackc/pgx/v5"
)

func (s *Store) CreateUser(ctx context.Context, email, passwordHash string) (user.User, error) {
	var u user.User

	err := s.pool.QueryRow(
		ctx,
		`INSERT INTO users (email, password_hash)
         VALUES ($1, $2)
         RETURNING id, email, password_hash, created_at`,
		email,
		passwordHash,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)

	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	var u user.User

	err := s.pool.QueryRow(
		ctx,
		`SELECT id, email, password_hash, created_at
		 FROM users
		 WHERE email = $1`,
		email,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return user.User{}, storage.ErrUserNotFound
	}

	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
