package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sohibjon7731/ecommerce_backend/config"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/handler"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	r := gin.Default()
	authHandler := handler.NewAuthHandler()
	api := r.Group("/api/v1/auth")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
	}

	r.Run(":8080")

}
