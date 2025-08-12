package http

import (
	"net/http"

	"github.com/Uikola/task-manager/internal/adapters/transport/http/v1/task"
)

func NewServer(taskHandler *task.Handler) http.Handler {
	router := http.NewServeMux()

	addRoutes(router, taskHandler)

	var handler http.Handler = router

	return handler
}

func addRoutes(router *http.ServeMux, taskHandler *task.Handler) {
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			taskHandler.Create(w, r)
		case http.MethodGet:
			taskHandler.List(w, r)
		}
	})
	router.HandleFunc("/tasks/{id}", taskHandler.GetByID)
}
