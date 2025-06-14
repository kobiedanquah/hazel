package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Workspace    *Workspace `json:"workspace,omitempty"`
	StartDate    time.Time  `json:"startDate,omitzero"`
	EndDate      time.Time  `json:"endDate,omitzero"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"createdAt"`
	LastModified time.Time  `json:"lastModified"`
}

type ProjectStore interface {
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, project *Project) error
	GetProject(ctx context.Context, id uuid.UUID) (*Project, error)
	GetWorkspaceProjects(ctx context.Context, workspaceId uuid.UUID) ([]Project, error)
	DeleteProject(ctx context.Context, id uuid.UUID) error
}
