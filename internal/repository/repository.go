package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Task
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Task: NewTaskRepository(db),
	}
}
