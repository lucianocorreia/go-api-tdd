package repository

import "database/sql"

// Repository is a struct that represents a repository.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new repository.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
