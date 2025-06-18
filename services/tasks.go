package services

import (
	"context"
	"time"

	"github.com/freekobie/hazel/models"
	"github.com/google/uuid"
)

func (s *WorkspaceService) CreateTask(ctx context.Context, task *models.Task) error {
	task.Id = uuid.New()

	lastModified := time.Now()
	task.CreatedAt = lastModified
	task.LastModified = lastModified
	task.Status = "todo"

	err := s.store.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (s *WorkspaceService) GetTask(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	return s.store.GetTask(ctx, id)
}

func (s *WorkspaceService) UpdateTask(ctx context.Context, data map[string]any) (*models.Task, error) {

	id, _ := data["id"]
	task, err := s.store.GetTask(ctx, id.(uuid.UUID))
	if err != nil {
		return nil, err
	}

	title, ok := data["title"]
	if ok {
		task.Title = title.(string)
	}

	description, ok := data["description"]
	if ok {
		task.Description = description.(string)
	}

	status, ok := data["status"]
	if ok {
		task.Status = models.TaskStatus(status.(string))
	}

	priority, ok := data["priority"]
	if ok {
		task.Priority = models.TaskPriority(priority.(string))
	}

	due, ok := data["due"]
	if ok {
		task.Due = due.(time.Time)
	}

	task.LastModified = time.Now()

	err = s.store.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *WorkspaceService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return s.store.DeleteTask(ctx, id)
}

func (s *WorkspaceService) GetProjectTasks(ctx context.Context, projectId uuid.UUID) ([]models.Task, error) {
	return s.store.GetTasksForProject(ctx, projectId)
}
