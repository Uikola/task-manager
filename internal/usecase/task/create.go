package task

import (
	"context"
	"time"

	"github.com/Uikola/task-manager/internal/adapters/transport/http/v1/dto/task/request"
	"github.com/Uikola/task-manager/internal/domain/entity"
)

func (uc *taskUsecase) Create(ctx context.Context, input request.CreateTask) (*entity.Task, error) {
	taskID, err := uc.uuidGenerator.GenerateV4()
	if err != nil {
		return nil, err
	}

	task := &entity.Task{
		ID:          taskID,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		CreatedAt:   time.Now(),
	}

	if err := uc.taskRepository.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}
