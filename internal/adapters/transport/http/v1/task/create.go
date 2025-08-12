package task

import (
	"encoding/json"
	"net/http"

	"github.com/Uikola/task-manager/internal/adapters/transport/http/v1/dto/task/request"
	"github.com/Uikola/task-manager/pkg/logger"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	h.asyncLogWriter.WriteLog(logger.InfoLevel, "new task creation started", nil, nil)

	var createTaskRequest request.CreateTask
	if err := json.NewDecoder(r.Body).Decode(&createTaskRequest); err != nil {
		h.asyncLogWriter.WriteLog(logger.DebugLevel, "failed to decode create task request", err, nil)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := createTaskRequest.Validate(); err != nil {
		h.asyncLogWriter.WriteLog(logger.DebugLevel, "invalid create task request", err, nil)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.taskUsecase.Create(ctx, createTaskRequest)
	if err != nil {
		h.asyncLogWriter.WriteLog(logger.ErrorLevel, "failed to create task", err, nil)
		http.Error(w, "failed to create task", http.StatusInternalServerError)
		return
	}

	h.asyncLogWriter.WriteLog(logger.InfoLevel, "new task created", nil, logger.Fields{
		"id":          task.ID,
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
		"create_at":   task.CreatedAt,
	})

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(task)
}
