package domain

import "errors"

var (
	// Users errors.
	ErrUserNotFound = errors.New("user not found")

	// Security errors.
	ErrInvalidKey   = errors.New("invalid key")
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)
