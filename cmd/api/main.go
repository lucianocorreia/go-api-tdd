package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/lucianocorreia/go-api-tdd/store/postgres"
)

const (
	postgresDns = "postgres://root:secret@localhost:5455/blog?sslmode=disable"
	dbDriver    = "postgres"
)

func main() {
	db, err := connectToDB(dbDriver)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	srv, err := setup(db)
	if err != nil {
		panic(err)
	}

	if err = srv.run(":3001"); err != nil {
		panic(err)
	}
}

func setup(db *sql.DB) (*server, error) {
	store := postgres.NewPostgresStore(db)

	srv := NewServer(store)

	srv.setupRoutes()

	return srv, nil
}

func connectToDB(driver string) (*sql.DB, error) {
	db, err := sql.Open(driver, postgresDns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err.Error())
	}

	return db, nil
}
