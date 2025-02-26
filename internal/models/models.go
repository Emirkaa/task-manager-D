package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

var newErr = errors.New("Title cant be empty")

type Task struct {
	ID        uint      `gorm:primaryKey`
	Title     string    `gorm:not null`
	Completed bool      `gorm:default:false`
	CreatedAt time.Time `gorm:autoCreateTime`
	UpdatedAt time.Time `gorm:autoUpdateTime`
}

func (t *Task) CreateBefore(tt *gorm.DB) (err error) {
	if t.Title == "" {
		return newErr
	}
	return nil

}
