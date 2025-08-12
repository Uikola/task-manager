package repository

import (
	"context"

	"github.com/Uikola/task-manager/internal/domain/entity"
)

// TaskRepository defines the interface for task data access operations.
type TaskRepository interface {
	// Create stores a new task.
	Create(ctx context.Context, task *entity.Task) error

	// GetByID retrieves a task by its unique id if found, otherwise returns entity.ErrTaskNotFound.
	GetByID(ctx context.Context, id string) (*entity.Task, error)

	// GetAllByStatuses retrieves all tasks, optionally filtered by status.
	GetAllByStatuses(ctx context.Context, statuses []entity.TaskStatus) ([]*entity.Task, error)
}
