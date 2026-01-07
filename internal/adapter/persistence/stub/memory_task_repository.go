package stub

import (
	"context"
	"errors"
	"sync"

	"github.com/Fumiya-Tahara/serverless-playground/internal/domain/model"
	"github.com/Fumiya-Tahara/serverless-playground/internal/domain/repository"
)

type memoryTaskRepository struct {
	mu    sync.RWMutex
	tasks map[string]*model.Task
}

func NewMemoryTaskRepository() repository.TaskRepository {
	tasks := make(map[string]*model.Task)

	t1, _ := model.NewTask("1", "Lambdaの学習", "API Gatewayとの連携を確認する")
	t2, _ := model.NewTask("2", "買い物", "卵，牛乳，小麦粉を買う")

	tasks[t1.ID()] = t1
	tasks[t2.ID()] = t2

	return &memoryTaskRepository{
		tasks: tasks,
	}
}

func (r *memoryTaskRepository) Save(ctx context.Context, task *model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID()] = task
	return nil
}

func (r *memoryTaskRepository) FindAll(ctx context.Context) ([]*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var res []*model.Task
	for _, t := range r.tasks {
		res = append(res, t)
	}
	return res, nil
}

func (r *memoryTaskRepository) FindByID(ctx context.Context, id string) (*model.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *memoryTaskRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tasks, id)
	return nil
}
