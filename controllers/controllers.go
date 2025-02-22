package controllers

import (
	"net/http"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var j = models.Task{}
	error1 := c.ShouldBindJSON(&j)
	if error1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error1.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "created", "goal": j})

}
