package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sohibjon7731/ecommerce_backend/config"
	authHandler "github.com/sohibjon7731/ecommerce_backend/internal/auth/handler"
	productHandler "github.com/sohibjon7731/ecommerce_backend/internal/product/handler"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	r := gin.Default()
	authHandler := authHandler.NewAuthHandler()
	productHandler:= productHandler.NewProductHandler()
	api := r.Group("/api/v1/auth")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.POST("/product/create", productHandler.Create)
		api.GET("/products", productHandler.GetAllProducts)
	}

	r.Run(":8080")

}
