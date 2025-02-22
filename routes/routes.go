package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {
	router.POST("/tasks", controllers.CreateTask())
	router.GET("/tasks", controllers.GetTask())
	router.PUT("/tasks/:id", controllers.PutTask())
	router.DELETE("/tasks/:id", controllers.DeleteTask())
}
