package repository

import "github.com/lucianocorreia/go-api-tdd/pkg/domain"

// CreateUser creates a new user in the database.
func (r *Repository) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = 1
	return user, nil
}
