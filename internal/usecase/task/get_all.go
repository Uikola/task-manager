package task

import (
	"context"

	"github.com/Uikola/task-manager/internal/domain/entity"
)

func (uc *taskUsecase) GetAll(ctx context.Context, statuses []entity.TaskStatus) ([]*entity.Task, error) {
	return uc.taskRepository.GetAllByStatuses(ctx, statuses)
}
