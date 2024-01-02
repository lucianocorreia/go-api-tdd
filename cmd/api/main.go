package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/lucianocorreia/go-api-tdd/pkg/security"
	"github.com/lucianocorreia/go-api-tdd/store/postgres"
)

const (
	postgresDns = "postgres://root:secret@localhost:5455/blog?sslmode=disable"
	dbDriver    = "postgres"
	key         = "110d9b962ee573804cced9930a469ea9f0b14ec0574f2c61957625634047bcd1"
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

	jwt, err := security.NewJWT(key)
	if err != nil {
		return nil, err
	}

	srv := NewServer(store, jwt)

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
