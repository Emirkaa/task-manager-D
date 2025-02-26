package main

import (
	"fmt"
	"log"
	"net/http"
	"taskmanager/internal/controllers"
	"taskmanager/internal/data_base/db"
	"taskmanager/internal/services"
	"taskmanager/internal/utills"
)

func main() {
	Logger := utills.NewLogger()

	db, err := db.InitDB()
	if err != nil {
		Logger.Fatal("Failed to connect to the db", err)
	}

	Tservice := services.NewService(db)

	Tcontroller := controllers.NewController(Tservice)

	http.Handler("/tasks", Tcontroller.GetTasks)
	http.Handler("/tasks/create", Tcontroller.CreateTask)

	port := "8080"
	fmt.Printf("Server is avalible on port %s/n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
