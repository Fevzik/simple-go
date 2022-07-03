package repository

import (
	"github.com/Fevzik/simple-go/internal/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

type Task interface {
	GetOneById(id int) (*domain.Task, error)
	GetList() ([]domain.Task, error)
}

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t TaskRepository) GetOneById(id int) (*domain.Task, error) {
	return &domain.Task{
		ID:        157,
		Title:     "some task by me",
		CreatedAt: time.Now(),
	}, nil
}

func (t TaskRepository) GetList() ([]domain.Task, error) {
	return []domain.Task{}, nil
}
