// internal/domain/repository/task_repository.go
package repository

import (
	"context"

	"github.com/Fumiya-Tahara/serverless-playground/internal/domain/model"
)

type TaskRepository interface {
	Save(ctx context.Context, task *model.Task) error
	FindAll(ctx context.Context) ([]*model.Task, error)
	FindByID(ctx context.Context, id string) (*model.Task, error)
	Delete(ctx context.Context, id string) error
}
