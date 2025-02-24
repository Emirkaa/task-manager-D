package controllers

import (
	"net/http"
	"taskmanager/models"
	"taskmanager/task"

	"net/http"

	"github.com/gin-gonic/gin"
)

var taskrepo = task.TaskRepository // Мне нужно инкапсулировать логику работы с данными, иначе будет непонятное очко :D

func CreateTask(c *gin.Context) {
	var j = models.Task{}
	error1 := c.ShouldBindJSON(&j)
	if error1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error1.Error()})
		return
	}

	if j.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Название задачи обязательно"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created", "goal": j})

}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	j, err := task.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := taskrepo.Delete("id")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func PutTask(c *gin.Context) {
	id := c.Param("id")
	err := taskrepo.Update("id")

}
