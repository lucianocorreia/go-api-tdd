package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	postgresDns = "postgres://root:secret@localhost:5455/blog?sslmode=disable"
	dbDriver    = "postgres"
)

func main() {
	_, err := connectToDB(dbDriver)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello, World!")
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
