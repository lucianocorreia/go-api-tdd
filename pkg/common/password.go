package common

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}

// CheckPassword checks if a password matches a hash
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
