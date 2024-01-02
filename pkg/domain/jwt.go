package domain

import (
	"time"
)

// JWTPayload represents the payload of a JWT token.
type JWTPayload struct {
	UserID   int64     `json:"user_id"`
	ExpireAt time.Time `json:"iat"`
	IssuedAt time.Time `json:"exp"`
	Token    string    `json:"token"`
}

// Valid validates the payload.
func (p *JWTPayload) Valid() error {
	if p.UserID == 0 {
		return ErrUserNotFound
	}

	if time.Now().After(p.ExpireAt) {
		return ErrTokenExpired
	}

	return nil
}
