package service

import "todo/pkg/repository"

type TaskService interface {
	Create(task repository.Task) error
	GetByID(id uint) (*repository.Task, error)
	GetAll() ([]repository.Task, error)
	Update(task repository.Task) error
	Delete(id uint) error
	GetPendingTasks() ([]repository.Task, error)
	GetCompletedTasks() ([]repository.Task, error)
	GetTasksByDate(date string, status string) ([]repository.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}