package inmemory

import (
	"context"
	"sync"

	"github.com/Uikola/task-manager/internal/adapters/repository"
	"github.com/Uikola/task-manager/internal/domain/entity"
)

// taskRepository implements the TaskRepository interface using in-memory storage
// It provides thread-safe operations for managing tasks
type taskRepository struct {
	mu    sync.RWMutex
	tasks map[string]*entity.Task
}

// NewTaskRepository creates and returns a new instance of the in-memory task repository.
func NewTaskRepository() repository.TaskRepository {
	return &taskRepository{
		tasks: make(map[string]*entity.Task),
	}
}

// Create stores a new task.
func (r *taskRepository) Create(_ context.Context, task *entity.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task

	return nil
}

// GetByID retrieves a task by its unique id if found, otherwise returns repository.ErrTaskNotFound.
func (r *taskRepository) GetByID(_ context.Context, id string) (*entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return nil, repository.ErrTaskNotFound
	}

	return task, nil
}

// GetAll retrieves all tasks, optionally filtered by status.
func (r *taskRepository) GetAll(_ context.Context, status *entity.TaskStatus) ([]*entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var filtered []*entity.Task
	for _, t := range r.tasks {
		if status == nil || t.Status == *status {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil
}
