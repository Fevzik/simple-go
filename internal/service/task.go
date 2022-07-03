package service

import (
	"github.com/Fevzik/simple-go/internal/domain"
	"github.com/Fevzik/simple-go/internal/repository"
)

type Task interface {
	GetOneById(id int) (*domain.Task, error)
	GetList() ([]domain.Task, error)
}

type TaskService struct {
	r repository.Task
}

func NewTaskService(r repository.Task) *TaskService {
	return &TaskService{r: r}
}

func (s TaskService) GetOneById(id int) (*domain.Task, error) {
	return s.r.GetOneById(id)
}

func (s TaskService) GetList() ([]domain.Task, error) {
	return s.r.GetList()
}
