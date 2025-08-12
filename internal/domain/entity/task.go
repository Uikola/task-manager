package entity

import (
	"errors"
	"time"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
	StatusFailed     TaskStatus = "failed"
)

func (s TaskStatus) Valid() bool {
	switch s {
	case StatusPending, StatusInProgress, StatusDone, StatusFailed:
		return true
	default:
		return false
	}
}

func ValidTaskStatus(status string) bool {
	switch status {
	case "pending", "in_progress", "done", "failed":
		return true
	default:
		return false
	}
}

func (s TaskStatus) String() string {
	return string(s)
}

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
}

// ErrTaskNotFound is returned when a task with the specified ID is not found.
var ErrTaskNotFound = errors.New("task not found")
