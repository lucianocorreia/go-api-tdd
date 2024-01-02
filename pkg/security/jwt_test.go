package security

import (
	"testing"
	"time"

	"github.com/lucianocorreia/go-api-tdd/pkg/common"
	"github.com/lucianocorreia/go-api-tdd/pkg/domain"
)

func TestJWTTokenWithShortKey(t *testing.T) {
	_, err := NewJWT("short-key")
	if err == nil {
		t.Errorf("expected error when key is short")
	}
}

func TestJWTTokenValid(t *testing.T) {
	key := common.RandomString(32)
	newJwt, err := NewJWT(key)
	if err != nil {
		t.Errorf("expected no error when key is long enough")
	}

	user := &domain.User{
		ID:    1,
		Name:  "test",
		Email: "teste@email.com",
	}

	tokenPayload, err := newJwt.CreateToken(user, 1*time.Minute)
	if err != nil {
		t.Errorf("expected no error when creating token: %v", err)
	}

	if tokenPayload.UserID != user.ID {
		t.Errorf("expected user id %d, got %d", user.ID, tokenPayload.UserID)
	}

	if tokenPayload.Token == "" {
		t.Errorf("expected token, got empty string")
	}

	if tokenPayload.ExpireAt.IsZero() {
		t.Errorf("expected expire at, got zero time")
	}

	_, err = newJwt.VerifyToken(tokenPayload.Token)
	if err != nil {
		t.Errorf("expected no error when verifying token: %v", err)
	}

	_, err = newJwt.VerifyToken(tokenPayload.Token + "invalid")
	if err == nil {
		t.Errorf("expected error when verifying invalid token")
	}
}

func TestJwtTokenExpired(t *testing.T) {
	key := common.RandomString(32)
	newJwt, err := NewJWT(key)
	if err != nil {
		t.Errorf("expected no error when key is long enough")
	}

	user := &domain.User{
		ID:    1,
		Name:  "test",
		Email: "teste@email.com",
	}

	expiredToken, err := newJwt.CreateToken(user, -1*time.Minute)
	if err != nil {
		t.Errorf("expected no error when creating token: %v", err)
	}

	_, err = newJwt.VerifyToken(expiredToken.Token)
	if err == nil {
		t.Errorf("expected error when verifying expired token")
	}

}
