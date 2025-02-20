package models

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Title   string `json:"title"`
	Descrip string `json:"description"`
	Status  string `json:"status"`
}
