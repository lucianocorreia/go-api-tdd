package main

import "github.com/gin-gonic/gin"

func (s *server) setupRoutes() {
	mux := gin.Default()

	v1 := mux.Group("/api/v1")

	v1.POST("/users/create", s.createUser)

	s.router = mux

}
