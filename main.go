package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"log"
	"taskmanager/models"
	"taskmanager/routes"

	_ "github.com/lib/pq"
)

var db *gorm.DB

func main() {

	var ERR error
	db, ERR = gorm.Open("postgres", "user=emir", "ssmode=disable", "dbname=taskmanager")
	if ERR != nil {
		log.Fatal("Error: ", ERR)
	}
	defer db.Close()

	db.AutoMigrate(&models.Task{})
	router := gin.Default()
	routes.SetupRouters(router)
	router.Run(":8080")

}
