package postgres

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	postgresDns = "postgres://root:secret@localhost:5455/blog?sslmode=disable"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	dbConn, err := sql.Open("postgres", postgresDns)
	if err != nil {
		panic(err)
	}

	if err = dbConn.Ping(); err != nil {
		panic(err)
	}

	testDB = dbConn

	code := m.Run()
	err = testDB.Close()
	if err != nil {
		panic(err)
	}

	os.Exit(code)
}
