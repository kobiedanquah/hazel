package services

import (
	"context"
	"time"

	"github.com/freekobie/hazel/models"
	"github.com/google/uuid"
)

func (s *WorkspaceService) CreateProject(ctx context.Context, project *models.Project) error {
	project.Id = uuid.New()

	lastModified := time.Now()
	project.CreatedAt = lastModified
	project.LastModified = lastModified
	project.Status = "active"

	err := s.store.CreateProject(ctx, project)
	if err != nil {
		return err
	}

	return nil
}

func (s *WorkspaceService) GetProject(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	return s.store.GetProject(ctx, id)
}
