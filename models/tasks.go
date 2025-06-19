package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TaskStatus defines allowed statuses for a task.
type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "started"
	StatusDone       TaskStatus = "complete"
)

// TaskPriority defines allowed priorities for a task.
type TaskPriority string

const (
	PriorityLow    TaskPriority = "low"
	PriorityMedium TaskPriority = "medium"
	PriorityHigh   TaskPriority = "high"
)

// Task represents a single work item within a project.
type Task struct {
	Id           uuid.UUID    `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Project      *Project     `json:"project,omitzero"`
	Status       TaskStatus   `json:"status"`
	Priority     TaskPriority `json:"priority"`
	Due          time.Time    `json:"due,omitzero"`
	CreatedAt    time.Time    `json:"createdAt"`
	LastModified time.Time    `json:"lastModified"`
}

type TaskStore interface {
	CreateTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, task *Task) error
	GetTask(ctx context.Context, id uuid.UUID) (*Task, error)
	GetTasksForProject(ctx context.Context, projectId uuid.UUID) ([]Task, error)
	DeleteTask(ctx context.Context, id uuid.UUID) error
	AssignTask(ctx context.Context, taskId, userId uuid.UUID) error
	UnassignTask(ctx context.Context, taskId, userId uuid.UUID) error
	GetAssignedUsers(ctx context.Context, taskId uuid.UUID) ([]User, error)
}
