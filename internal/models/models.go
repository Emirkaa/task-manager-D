package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type Task struct {
	ID        uint      `gorm:primaryKey`
	Title     string    `gorm:not null`
	Completed bool      `gorm:default false`
	CreatedAt time.Time `gorm:autoCreateTime`
	UpdatedAt time.Time `gorm:autoUpdateTime`
}

func (t *Task) CreateBefore(tt *gorm.DB) (err error) {
	if t.Title == "" {
		return fmt.Errorf("Empty title!")
	}
	return nil

}
