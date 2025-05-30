package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id          uuid.UUID  `json:"id"`
	WorkspaceID uuid.UUID  `json:"workspaceId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	OwnerID     uuid.UUID  `json:"ownerId"`
	Owner       *User      `json:"owner,omitempty"`
	Workspace   *Workspace `json:"workspace,omitempty"`
	StartDate   time.Time  `json:"startDate,omitzero"`
	EndDate     time.Time  `json:"endDate,omitzero"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type ProjectStore interface {
	Create(ctx context.Context, project *Project) error
	Update(ctx context.Context, project *Project) error
	Get(ctx context.Context, id uuid.UUID) (Project, error)
	GetAllForWorkspace(ctx context.Context, workspaceId uuid.UUID) ([]Project, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
