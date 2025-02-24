package task

import (
	"taskmanager/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) (*models.Task, error)
	GetByID(id string) (*models.Task, error)
	Update(id string) (*models.Task, error)
	Delete(id string) (*models.Task, error)
}

type taskrepository struct {
	db *gorm.DB
}

func (t *taskrepository) Create(task *models.Task) (*models.Task, error) {
	if err := t.db.Create(task).Error; err != nil {
		return err, nil
	}
	return task, nil
}
