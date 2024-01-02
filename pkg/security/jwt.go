package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lucianocorreia/go-api-tdd/pkg/domain"
)

type jwtToken struct {
	key string
}

// NewJWT creates a new JWT token.
func NewJWT(key string) (domain.JWT, error) {
	if len(key) < 32 {
		return nil, domain.ErrInvalidKey
	}

	return &jwtToken{key}, nil
}

// CreateToken creates a new token for the given user.
func (j *jwtToken) CreateToken(user *domain.User, duration time.Duration) (*domain.JWTPayload, error) {
	now := time.Now()

	payload := &domain.JWTPayload{
		UserID:   user.ID,
		ExpireAt: now.Add(duration),
		IssuedAt: now,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(j.key))
	if err != nil {
		return nil, err
	}

	payload.Token = tokenString

	return payload, nil
}

// VerifyToken verifies the given token.
func (j *jwtToken) VerifyToken(tokenString string) (*domain.JWTPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, domain.ErrInvalidToken
		}
		return []byte(j.key), nil
	})

	if err != nil {
		var validationError *jwt.ValidationError
		ok := errors.As(err, &validationError)
		if ok && errors.Is(validationError.Inner, domain.ErrTokenExpired) {
			return nil, domain.ErrTokenExpired
		}

		return nil, domain.ErrInvalidToken
	}

	payload, ok := token.Claims.(*domain.JWTPayload)
	if !ok {
		return nil, domain.ErrInvalidToken
	}

	payload.Token = tokenString

	return payload, nil
}
