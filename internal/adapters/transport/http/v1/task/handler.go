package task

import (
	"github.com/Uikola/task-manager/internal/adapters/logwriter"
	"github.com/Uikola/task-manager/internal/usecase"
)

type Handler struct {
	asyncLogWriter logwriter.LogWriter

	taskUsecase usecase.TaskUsecase
}

func NewHandler(asyncLogWriter logwriter.LogWriter, taskUsecase usecase.TaskUsecase) *Handler {
	return &Handler{
		asyncLogWriter: asyncLogWriter,

		taskUsecase: taskUsecase,
	}
}
