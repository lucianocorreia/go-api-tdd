package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucianocorreia/go-api-tdd/pkg/domain"
	"github.com/lucianocorreia/go-api-tdd/store/postgres"
)

var (
	testStore domain.Store
)

type testServer struct {
	*httptest.Server
}

func newTestServer(h http.Handler) *testServer {
	return &testServer{
		Server: httptest.NewServer(h),
	}
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	db, err := connectToDB(dbDriver)
	if err != nil {
		panic(err)
	}

	testStore = postgres.NewPostgresStore(db)

	code := m.Run()

	// Clean up database.
	_ = testStore.DeleteAllUsers()

	os.Exit(code)
}
