package usecase

import (
	"context"

	"github.com/Uikola/task-manager/internal/adapters/transport/http/v1/dto/task/request"
	"github.com/Uikola/task-manager/internal/domain/entity"
)

type TaskUsecase interface {
	Create(ctx context.Context, input request.CreateTask) (*entity.Task, error)

	GetByID(ctx context.Context, id string) (*entity.Task, error)

	GetAll(ctx context.Context, status *entity.TaskStatus) ([]*entity.Task, error)
}
