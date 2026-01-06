package task

import "time"

type CreateTaskInput struct {
	Title   string
	Content string
}

type UpdateTaskInput struct {
	ID      string
	Title   string
	Content string
}

type TaskOutput struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
