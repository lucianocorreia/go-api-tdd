package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucianocorreia/go-api-tdd/pkg/domain"
)

type server struct {
	router *gin.Engine
	store  domain.Store
}

// NewServer creates a new server instance.
func NewServer(store domain.Store) *server {
	return &server{
		store: store,
	}
}

func (s *server) routes() *gin.Engine {
	if s.router == nil {
		s.setupRoutes()
	}

	return s.router
}

func (s *server) run(addr string) error {
	return s.routes().Run(addr)
}
