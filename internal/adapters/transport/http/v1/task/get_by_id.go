package task

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Uikola/task-manager/internal/domain/entity"
	"github.com/Uikola/task-manager/pkg/logger"
)

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	id := r.PathValue("id")
	h.asyncLogWriter.WriteLog(logger.InfoLevel, "get task by id started", nil, logger.Fields{
		"id": id,
	})

	task, err := h.taskUsecase.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrTaskNotFound):
			h.asyncLogWriter.WriteLog(logger.DebugLevel, "task not found", nil, logger.Fields{
				"id": id,
			})
			http.Error(w, "task with this id not found", http.StatusNotFound)
			return
		case err != nil:
			h.asyncLogWriter.WriteLog(logger.ErrorLevel, "failed to get task by id", err, logger.Fields{
				"id": id,
			})
			http.Error(w, "failed to get task by id", http.StatusInternalServerError)
			return
		}
	}

	h.asyncLogWriter.WriteLog(logger.InfoLevel, "got task by id", nil, logger.Fields{
		"id": id,
	})

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(task)
}
