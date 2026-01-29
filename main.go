package main

import (
	"user-auth/config"
	"user-auth/handler"
	"user-auth/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	r := gin.Default()
	r.Use(middleware.RecoveryMiddleware())

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	r.Run(":8080")
}
