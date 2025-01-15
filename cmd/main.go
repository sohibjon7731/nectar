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
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	authHandler := authHandler.NewAuthHandler()
	productHandler:= productHandler.NewProductHandler()
	api := r.Group("/api/v1")
	{
		api.POST("/auth/register", authHandler.Register)
		api.POST("/auth/login", authHandler.Login)
		api.POST("/product/create", productHandler.Create)
		api.GET("/products", productHandler.GetAllProducts)
	}

	r.Run(":8080")

}
