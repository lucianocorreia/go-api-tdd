package postgres

import (
	"errors"
	"fmt"
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

	store := NewPostgresStore(testDB)

	createdUser, err := store.CreateUser(user)
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

	// Delete user from database.
	err = store.DeleteUserByID(createdUser.ID)
	if err != nil {
		t.Fatalf("error deleting user: %s", err.Error())
	}

}

func TestFindUserByEmail(t *testing.T) {
	store := NewPostgresStore(testDB)

	user := domain.User{
		Name:     "Jown Doe",
		Email:    "teste@email.com",
		Password: "password",
	}

	createdUser, err := store.CreateUser(&user)
	if err != nil {
		t.Fatalf("error creating user: %s", err.Error())
	}

	foundUser, err := store.FindUserByEmail(createdUser.Email)
	if err != nil {
		t.Fatalf("error finding user: %s", err.Error())
	}

	if foundUser.ID != createdUser.ID {
		t.Errorf("expected user id to be %d got %d", createdUser.ID, foundUser.ID)
	}

	_, err = store.FindUserByEmail("invalid@email.com")
	if err == nil {
		t.Fatalf("expected error finding user")
	}

	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		t.Fatalf("expected error to be ErrUserNotFound got %s", err.Error())
	}

	// Delete user from database.
	err = store.DeleteUserByID(createdUser.ID)
	if err != nil {
		t.Fatalf("error deleting user: %s", err.Error())
	}
}

func TestFindUserByID(t *testing.T) {
	store := NewPostgresStore(testDB)

	user := domain.User{
		Name:     "Jown Doe",
		Email:    "teste@email.com",
		Password: "password",
	}

	createdUser, err := store.CreateUser(&user)
	if err != nil {
		t.Fatalf("error creating user: %s", err.Error())
	}

	fmt.Println(createdUser.ID)

	foundUser, err := store.FindUserByID(createdUser.ID)
	if err != nil {
		t.Fatalf("error finding user: %s", err.Error())
	}

	if foundUser.ID != createdUser.ID {
		t.Errorf("expected user id to be %d got %d", createdUser.ID, foundUser.ID)
	}

	_, err = store.FindUserByID(-987)
	if err == nil {
		t.Fatalf("expected error finding user")
	}

	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		t.Fatalf("expected error to be ErrUserNotFound got %s", err.Error())
	}

	// Delete user from database.
	err = store.DeleteUserByID(createdUser.ID)
	if err != nil {
		t.Fatalf("error deleting user: %s", err.Error())
	}
}
