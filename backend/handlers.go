package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/todos
func GetTodos(c *gin.Context) {
	var todos []Todo
	DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// GET /api/todos/:id
func GetTodo(c *gin.Context) {
	var todo Todo
	if err := DB.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// POST /api/todos
func CreateTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// PUT /api/todos/:id
func UpdateTodo(c *gin.Context) {
	var todo Todo
	if err := DB.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&todo).Updates(input)
	c.JSON(http.StatusOK, todo)
}

// DELETE /api/todos/:id
func DeleteTodo(c *gin.Context) {
	var todo Todo
	if err := DB.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
