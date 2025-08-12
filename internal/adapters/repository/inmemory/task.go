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

// GetByID retrieves a task by its unique id if found, otherwise returns entity.ErrTaskNotFound.
func (r *taskRepository) GetByID(_ context.Context, id string) (*entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		return nil, entity.ErrTaskNotFound
	}

	return task, nil
}

// GetAllByStatuses retrieves all tasks, optionally filtered by status.
func (r *taskRepository) GetAllByStatuses(ctx context.Context, statuses []entity.TaskStatus) ([]*entity.Task, error) {
	if len(statuses) == 0 {
		return r.getAll(ctx)
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	statusSet := make(map[string]bool)
	for _, status := range statuses {
		statusSet[status.String()] = true
	}

	var filteredTasks []*entity.Task
	for _, task := range r.tasks {
		if statusSet[task.Status.String()] {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}

// getAll retrieves all tasks
func (r *taskRepository) getAll(_ context.Context) ([]*entity.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*entity.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
