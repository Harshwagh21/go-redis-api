package cache

import (
	"net/http"

	"github.com/Harshwagh21/go-redis-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Set(c *gin.Context) {
	var req domain.CacheRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Set(c.Request.Context(), req.Key, req.Value, req.TTL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "key set Successfully",
		"key":     req.Key,
		"ttl":     req.TTL,
	})
}

func (h *Handler) Get(c *gin.Context) {
	key := c.Param("key")

	value, err := h.service.Get(c.Request.Context(), key)
	if err != nil {
		if err.Error() == "key not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"key":   key,
		"value": value,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	key := c.Param("key")

	if err := h.service.Delete(c.Request.Context(), key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "key deleted sucessfully",
		"key":     key,
	})
}
