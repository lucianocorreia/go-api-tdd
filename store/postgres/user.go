package postgres

import (
	"database/sql"
	"errors"

	"github.com/lucianocorreia/go-api-tdd/pkg/common"
	"github.com/lucianocorreia/go-api-tdd/pkg/domain"
)

const (
	sqlCreateUser      = "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, password"
	sqlDeleteUserByID  = "DELETE FROM users WHERE id = $1"
	sqlFindUserByEmail = "SELECT id, name, email, password FROM users WHERE email = $1"
	sqlFindUserByID    = "SELECT id, name, email, password FROM users WHERE id = $1"
)

// CreateUser creates a new user in the database.
func (s *postgresStore) CreateUser(user *domain.User) (*domain.User, error) {
	hashedPassword, err := common.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	err = s.db.QueryRow(sqlCreateUser, user.Name, user.Email, user.Password).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUserByID deletes a user from the database by its ID.
func (s *postgresStore) DeleteUserByID(id int64) error {
	_, err := s.db.Exec(sqlDeleteUserByID, id)
	if err != nil {
		return err
	}

	return nil
}

// FindUserByEmail finds a user from the database by its email.
func (s *postgresStore) FindUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	err := s.db.QueryRow(sqlFindUserByEmail, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

// FindUserByID finds a user from the database by its ID.
func (s *postgresStore) FindUserByID(ID int64) (*domain.User, error) {
	user := &domain.User{}
	err := s.db.QueryRow(sqlFindUserByID, ID).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *postgresStore) DeleteAllUsers() error {
	_, err := s.db.Exec("DELETE FROM users")
	if err != nil {
		return err
	}

	return nil
}
