package main

import (
	userService "example/context/hello/internal/core/services/user"
	"example/context/hello/internal/handlers/userHandler"
	"example/context/hello/internal/repositories/userRepositories"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepository := userRepositories.NewMemKVS()
	userService := userService.New(userRepository)
	userHandler := userHandler.NewHTTPHandler(userService)

	router := gin.New()
	router.GET("/user/:id", userHandler.Get)
	router.POST("/user", userHandler.Create)

	router.Run(":8080")
}