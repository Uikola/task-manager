package task

import (
	"context"

	"github.com/Uikola/task-manager/internal/domain/entity"
)

func (uc *taskUsecase) GetByID(ctx context.Context, id string) (*entity.Task, error) {
	return uc.taskRepository.GetByID(ctx, id)
}
