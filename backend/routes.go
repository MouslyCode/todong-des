package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/todos", GetTodos)
		api.GET("/todos/:id", GetTodo)
		api.POST("/todos", CreateTodo)
		api.PUT("/todos/:id", UpdateTodo)
		api.DELETE("/todos/:id", DeleteTodo)
	}
}
