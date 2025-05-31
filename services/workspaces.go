package services

import (
	"context"

	"github.com/freekobie/hazel/models"
)

type WorkspaceService struct {
	store models.WorkspaceStore
}

func NewWorkspaceService(store models.WorkspaceStore) *WorkspaceService {
	return &WorkspaceService{
		store: store,
	}
}

func (ws *WorkspaceService) NewWorkspace(ctx context.Context, w *models.Workspace) error {

	return nil
}
