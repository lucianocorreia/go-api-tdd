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
