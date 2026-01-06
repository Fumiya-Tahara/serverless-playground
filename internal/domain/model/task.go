package model

import (
	"errors"
	"time"
)

type Task struct {
	id        string
	title     string
	content   string
	createdAt time.Time
	updatedAt time.Time
}

func NewTask(id, title, content string) (*Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	now := time.Now()
	return &Task{
		id:        id,
		title:     title,
		content:   content,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func (t *Task) Update(title, content string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}

	t.title = title
	t.content = content
	t.updatedAt = time.Now()
	return nil
}

func (t *Task) ID() string           { return t.id }
func (t *Task) Title() string        { return t.title }
func (t *Task) Content() string      { return t.content }
func (t *Task) CreatedAt() time.Time { return t.createdAt }
func (t *Task) UpdatedAt() time.Time { return t.updatedAt }

func ReconstructTask(id, title, content string, createdAt, updatedAt time.Time) *Task {
	return &Task{
		id:        id,
		title:     title,
		content:   content,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
