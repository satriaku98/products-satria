package main

import (
	"erajaya/config"
	"erajaya/internal/middleware"
	"erajaya/internal/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "erajaya/docs"
)

// @title Erajaya API Documentation
// @version 1.0
// @description This is an API documentation for Erajaya.
// @BasePath /api/v1
func main() {
	// Env
	config.LoadEnv()

	// Logger
	logger := config.NewLogger()

	// Database
	db := config.NewDB(logger)
	rdb := config.NewRedis(logger)

	// Gin
	r := gin.New()

	// Middleware
	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(gin.Recovery())
	r.Use(middleware.RateLimiterMiddleware())

	// Routes
	routes.RegisterProductRoutes(r, db, rdb)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
