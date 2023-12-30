package postgres

import "database/sql"

// postgresStore represents a PostgreSQL implementation of the store.
type postgresStore struct {
	db *sql.DB
}

// NewPostgresStore creates a new PostgreSQL store.
func NewPostgresStore(db *sql.DB) *postgresStore {
	return &postgresStore{
		db: db,
	}
}
