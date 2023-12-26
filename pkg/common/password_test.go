package common

import "testing"

func TestPasswordHash(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Fatalf("error hashing password: %s", err.Error())
	}

	if password == hashedPassword {
		t.Errorf("expected hashed password to be different from plain password")
	}

	if len(hashedPassword) == 0 {
		t.Errorf("expected hashed password to be set")
	}
}

func TestPasswordHashError(t *testing.T) {
	longPassword := make([]byte, 73)
	_, err := HashPassword(string(longPassword))
	if err == nil {
		t.Errorf("expected error hashing password")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("error hashing password: %s", err.Error())
	}

	err = CheckPassword(password, hashedPassword)
	if err != nil {
		t.Errorf("expected password to be valid")
	}
}
