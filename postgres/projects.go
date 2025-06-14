package postgres

import (
	"context"

	"github.com/freekobie/hazel/models"
	"github.com/google/uuid"
)

func (w *WorkspaceStore) CreateProject(ctx context.Context, project *models.Project) error {

	return nil

}

func (w *WorkspaceStore) UpdateProject(ctx context.Context, project *models.Project) error {
	return nil
}

func (w *WorkspaceStore) GetProject(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	return nil, nil
}

func (w *WorkspaceStore) GetWorkspaceProjects(ctx context.Context, workspaceId uuid.UUID) ([]models.Project, error) {
	return nil, nil
}

func (w *WorkspaceStore) DeleteProject(ctx context.Context, id uuid.UUID) error {
	return nil
}
