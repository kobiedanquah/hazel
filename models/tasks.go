package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TaskStatus defines allowed statuses for a task.
type TaskStatus string

const (
	StatusTodo       TaskStatus = "To Do"
	StatusInProgress TaskStatus = "In Progress"
	StatusDone       TaskStatus = "Done"
)

// TaskPriority defines allowed priorities for a task.
type TaskPriority string

const (
	PriorityLow    TaskPriority = "Low"
	PriorityMedium TaskPriority = "Medium"
	PriorityHigh   TaskPriority = "High"
)

// Task represents a single work item within a project.
type Task struct {
	Id          uuid.UUID    `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	ProjectID   uuid.UUID    `json:"projectId"`
	AssigneeID  uuid.UUID    `json:"assigneeId,omitzero"`
	Assignee    *User        `json:"assignee,omitempty"`
	ReporterID  uuid.UUID    `json:"reporterId"`
	Reporter    *User        `json:"reporter,omitempty"`
	Status      TaskStatus   `json:"status"`
	Priority    TaskPriority `json:"priority"`
	DueDate     time.Time    `json:"dueDate,omitzero"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

type TaskStore interface {
	Create(ctx context.Context, task *Task) error
	Update(ctx context.Context, task *Task) error
	Get(ctx context.Context, id uuid.UUID) (Task, error)
	GetAllForProject(ctx context.Context, projectId uuid.UUID)
	Delete(ctx context.Context, id uuid.UUID) error
}
