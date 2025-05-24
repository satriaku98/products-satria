package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"erajaya/internal/handler"
	"erajaya/internal/repository"
	"erajaya/internal/service"
)

func RegisterProductRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	// Dependency injection
	productRepo := repository.NewProductRepository(db, rdb)
	productService := service.NewService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Routes
	productGroup := r.Group("api/v1/products")
	{
		productGroup.POST("", productHandler.AddProduct)
		productGroup.GET("", productHandler.ListProduct)
	}
}
