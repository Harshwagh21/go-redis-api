package health

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	redis *redis.Client
}

func NewHandler(redis *redis.Client) *Handler {
	return &Handler{redis: redis}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	start := time.Now()

	err := h.redis.Ping(context.Background()).Err()

	latency := time.Since(start).Milliseconds()

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "healthy",
		"redis":       "connected",
		"latency_ms":  latency,
		"server_time": time.Now().UTC().Format(time.RFC3339),
	})
}
