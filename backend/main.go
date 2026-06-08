package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Connect()

	r := gin.Default()

	// CORS middleware - biar frontend bisa akses API
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	SetupRoutes(r)
	r.Run(":8080")
}
