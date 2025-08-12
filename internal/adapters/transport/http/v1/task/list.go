package task

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Uikola/task-manager/internal/domain/entity"

	"github.com/Uikola/task-manager/pkg/logger"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	query := r.URL.Query()

	statusesParam := query.Get("statuses")

	var statuses []entity.TaskStatus
	if statusesParam != "" {
		statuses = make([]entity.TaskStatus, 0)
		params := strings.Split(statusesParam, ",")
		for _, param := range params {
			temp := strings.TrimSpace(param)
			if !entity.ValidTaskStatus(temp) {
				continue
			}

			statuses = append(statuses, entity.TaskStatus(temp))
		}
	}

	h.asyncLogWriter.WriteLog(logger.InfoLevel, "list task started", nil, logger.Fields{
		"statuses_to_filter": statuses,
	})

	tasks, err := h.taskUsecase.GetAll(ctx, statuses)
	if err != nil {
		h.asyncLogWriter.WriteLog(logger.ErrorLevel, "failed to list tasks", err, nil)
		http.Error(w, "failed to", http.StatusInternalServerError)
		return
	}

	h.asyncLogWriter.WriteLog(logger.InfoLevel, "list task", nil, logger.Fields{
		"tasks_count": len(tasks),
	})

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}
