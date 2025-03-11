package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"taskmanager/internal/controllers"
	"taskmanager/internal/database"
	"taskmanager/internal/services"
	"taskmanager/internal/utills"
)

func main() {
	logger := utills.NewLogger()

	db, err := database.InitDB()
	if err != nil {
		logger.Fatal("Failed to connect to the db", err)
	}

	Tservice := services.NewService(db)

	Tcontroller := controllers.NewController(Tservice)

	http.HandleFunc("/tasks", Tcontroller.GetTasks)
	http.HandleFunc("/tasks/create", Tcontroller.CreateTask)

	port := "8080"
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	go func() {
		logger.Infof("Server if available on port %s\n", port)
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal("Server Failed: ", err)
		}

	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q

	err = srv.Shutdown(context.Background()) // можно nil.
	if err != nil {
		logger.Fatal("Server disable error")

	}
	logger.Info("Server disable")

}
