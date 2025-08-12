package task

import (
	"context"

	"github.com/Uikola/task-manager/internal/domain/entity"
)

func (uc *taskUsecase) GetAll(ctx context.Context, status *entity.TaskStatus) ([]*entity.Task, error) {
	return uc.taskRepository.GetAll(ctx, status)
}
