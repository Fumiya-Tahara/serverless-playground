package task

import (
	"context"
	"fmt"

	"github.com/Fumiya-Tahara/serverless-playground/internal/domain/model"
	"github.com/Fumiya-Tahara/serverless-playground/internal/domain/repository"
)

type TaskUsecase interface {
	Create(ctx context.Context, in CreateTaskInput) error
	FindAll(ctx context.Context) ([]TaskOutput, error)
	Update(ctx context.Context, in UpdateTaskInput) error
	Delete(ctx context.Context, id string) error
}

type taskInteractor struct {
	repo repository.TaskRepository
}

func NewTaskInteractor(r repository.TaskRepository) TaskUsecase {
	return &taskInteractor{repo: r}
}

func (i *taskInteractor) Create(ctx context.Context, in CreateTaskInput) error {
	id := "generated-uuid"

	task, err := model.NewTask(id, in.Title, in.Content)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return i.repo.Save(ctx, task)
}

func (i *taskInteractor) FindAll(ctx context.Context) ([]TaskOutput, error) {
	tasks, err := i.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	outputs := make([]TaskOutput, len(tasks))
	for j, t := range tasks {
		outputs[j] = TaskOutput{
			ID:        t.ID(),
			Title:     t.Title(),
			Content:   t.Content(),
			CreatedAt: t.CreatedAt(),
		}
	}
	return outputs, nil
}

func (i *taskInteractor) Update(ctx context.Context, in UpdateTaskInput) error {
	task, err := i.repo.FindByID(ctx, in.ID)
	if err != nil {
		return err
	}

	if err := task.Update(in.Title, in.Content); err != nil {
		return err
	}

	return i.repo.Save(ctx, task)
}

func (i *taskInteractor) Delete(ctx context.Context, id string) error {
	return i.repo.Delete(ctx, id)
}
