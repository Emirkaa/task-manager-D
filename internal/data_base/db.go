package db

import (
	"fmt"
	"log"
	"taskmanager/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=Emir user=Emir password=password db_name=task_manager port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)

	}
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("Failed to migrate: ", err)
	}

	DB = db
	fmt.Println("Successful to connect to db")

}
