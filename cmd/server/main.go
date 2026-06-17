package main

import (
	"github.com/Epimetheus29/go-sandbox/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	api := r.Group("/api")
	{
		api.GET("/flip", handlers.FlipCasing)
		api.GET("/fibonacci", handlers.Fibonacci)
	}

	r.Run(":8080")
}
