package repository

import "gorm.io/gorm"

type TaskRepository interface {
	Create(task Task) error
	GetByID(id uint) (*Task, error)
	GetAll() ([]Task, error)
	Update(task Task) error
	Delete(id uint) error
	GetPendingTasks() ([]Task, error)
	GetCompletedTasks() ([]Task, error)
	GetTasksByDate(date string, status string) ([]Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}
