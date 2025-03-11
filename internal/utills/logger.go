package utills

import (
	"fmt"
	"log"
	"os"
)

func NewLogger() *log.Logger {
	file, err := os.OpenFile("task.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file log: ", err)
		os.Exit(1)

	}
	logger := log.New(file, "TASK_LOG", log.Ldate|log.Ltime|log.Lshortfile)
	return logger

}
