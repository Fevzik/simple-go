package service

import "github.com/Fevzik/simple-go/internal/repository"

type Service struct {
	Task
}

func NewService(r repository.Repository) *Service {
	return &Service{
		Task: NewTaskService(r.Task),
	}
}
