package service

import "errors"

var (
	ErrInvalidPassword    = errors.New("password hash failed")
	ErrEmailAlreadyExists = errors.New("email already exists")
)
