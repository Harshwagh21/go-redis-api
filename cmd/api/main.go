package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Harshwagh21/go-redis-api/internal/cache"
	"github.com/Harshwagh21/go-redis-api/internal/health"
	"github.com/Harshwagh21/go-redis-api/pkg/redisclient"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	redisclient.Init()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)

	healthHandler := health.NewHandler(redisclient.Client)

	cacheRepo := cache.NewRepository(redisclient.Client)
	cacheService := cache.NewService(cacheRepo)
	cacheHandler := cache.NewHandler(cacheService)

	r.GET("/health", healthHandler.HealthCheck)

	r.POST("/cache", cacheHandler.Set)
	r.GET("/cache/:key", cacheHandler.Get)
	r.DELETE("/cache/:key", cacheHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
