package domain

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}
