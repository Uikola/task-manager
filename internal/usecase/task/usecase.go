package task

import (
	"github.com/Uikola/task-manager/internal/adapters/repository"
	"github.com/Uikola/task-manager/internal/usecase"
	"github.com/Uikola/task-manager/pkg/uuid"
)

type taskUsecase struct {
	taskRepository repository.TaskRepository

	uuidGenerator uuid.Generator
}

func NewUsecase(taskRepository repository.TaskRepository, uuidGenerator uuid.Generator) usecase.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,

		uuidGenerator: uuidGenerator,
	}
}
