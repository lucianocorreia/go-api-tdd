package repository

import (
	"testing"

	"github.com/lucianocorreia/go-api-tdd/pkg/domain"
)

func TestCreateUser(t *testing.T) {
	plainPassword := "password"
	user := &domain.User{
		Name:     "Jown Doe",
		Email:    "test@test.com",
		Password: plainPassword,
	}

	repo := NewRepository(nil)

	createdUser, err := repo.CreateUser(user)
	if err != nil {
		t.Fatalf("error creating user: %s", err.Error())
	}

	if createdUser.ID == 0 {
		t.Errorf("expected user id to be set got %d", createdUser.ID)
	}

	if user.Name != createdUser.Name {
		t.Errorf("expected user name to be set got %s", createdUser.Name)
	}

	if user.Email != createdUser.Email {
		t.Errorf("expected user email to be set got %s", createdUser.Email)
	}

	if plainPassword == createdUser.Password {
		t.Errorf("expected user password to be hashed got %s", createdUser.Password)
	}

}
