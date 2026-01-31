package entity

import "time"

type Task struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	IsCompleted bool       `json:"is_completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type UpdateTaskPayload struct {
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}
