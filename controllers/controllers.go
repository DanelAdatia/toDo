package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/praveeenrm/todoapp/config"
	"github.com/praveeenrm/todoapp/models"
)

// GetAllTodosHandler gets all the todos
func GetAllTodosHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET")
	var todos []models.Todo
	err := config.DB.Find(&todos).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

// GetATodoHandler gets a particular todo with id
func GetATodoHandler(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := config.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

// CreateATodoHandler creates a new todo
func CreateATodoHandler(c *gin.Context) {

	var todo models.Todo
	c.BindJSON(&todo)
	err := config.DB.Create(&todo).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// UpdateATodoHandler updates an existing todo
func UpdateATodoHandler(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	// get todo
	err := config.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.BindJSON(&todo)
	// update
	err = config.DB.Save(&todo).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

// DeleteATodoHandler delete/remove a particular todo
func DeleteATodoHandler(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := config.DB.Where("id = ?", id).Delete(&todo).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully deleted",
		})
	}
}
