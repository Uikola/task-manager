package repository

import (
	"context"
	"errors"

	"github.com/Uikola/task-manager/internal/domain/entity"
)

// ErrTaskNotFound is returned when a task with the specified ID is not found.
var ErrTaskNotFound = errors.New("task not found")

// TaskRepository defines the interface for task data access operations.
type TaskRepository interface {
	// Create stores a new task.
	Create(ctx context.Context, task *entity.Task) error

	// GetByID retrieves a task by its unique id if found, otherwise returns ErrTaskNotFound.
	GetByID(ctx context.Context, id string) (*entity.Task, error)

	// GetAll retrieves all tasks, optionally filtered by status.
	GetAll(ctx context.Context, status *entity.TaskStatus) ([]*entity.Task, error)
}
