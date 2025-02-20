package main

import "github.com/gin-gonic/gin"
import "github.com/jinzhu/gorm"
import _"github.com/lib/pq"
import "github.com/log"



var db *gorm.DB
func main(){


	var ERR error
	db, ERR = gorm.Open("postgres","user=emir","ssmode=disable","dbname=taskmanager")
	if ERR != nil{
		log.Fatal("Error: ",ERR)
	}
	defer db.close
	
	db.AutoMigrate(&models.Task{})
	router := gin.Default()
	SetupRouter(router)
	router.Run(":8080")

	
}
