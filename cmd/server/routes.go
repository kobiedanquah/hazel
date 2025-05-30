package main

import (
	"github.com/gin-gonic/gin"
)

func (s *application) routes() *gin.Engine {
	router := gin.Default()

	// users
	router.POST("/users", s.handler.CreateUser)
	router.POST("/users/verify", s.handler.VerifyUser)
	router.GET("/users/:id", s.handler.GetUser)
	router.DELETE("/users/:id", s.handler.DeleteUser)

	return router
}
