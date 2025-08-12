package request

import (
	"fmt"

	"github.com/Uikola/task-manager/internal/domain/entity"
)

type CreateTask struct {
	Title       string            `json:"title"`
	Description string            `json:"description,omitempty"`
	Status      entity.TaskStatus `json:"status"`
}

func (r *CreateTask) Validate() error {
	if !r.Status.Valid() {
		return fmt.Errorf("task status is invalid")
	}

	return nil
}
