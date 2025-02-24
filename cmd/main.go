package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/sohibjon7731/ecommerce_backend/cmd/docs"

	docs "github.com/sohibjon7731/ecommerce_backend/cmd/docs"
	"github.com/sohibjon7731/ecommerce_backend/config"
	authHandler "github.com/sohibjon7731/ecommerce_backend/internal/auth/handler"
	productHandler "github.com/sohibjon7731/ecommerce_backend/internal/product/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for Swagger integration.
// @BasePath /api/v1
func main() {
	if err := config.LoadConfig(); err != nil {
		log.Println("Warning: Config file not found or invalid:", err)
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	authH := authHandler.NewAuthHandler()
	productH := productHandler.NewProductHandler()

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authH.Register)
			auth.POST("/login", authH.Login)
		}

		products := api.Group("/products")
		{
			products.POST("/create", productH.Create)
			products.GET("/all", productH.GetAllProducts)
			products.PUT("/update/:id", productH.UpdateProduct)
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
