package service

import (
	"context"

	"github.com/nextri/product-road/db"
	"github.com/nextri/product-road/model"
)

// ProjectService handles business logic for projects.
type ProjectService struct {
	repo db.ProjectRepository
}

// NewUserService creates a new instance of ProjectService.
func NewProjectService(repo db.ProjectRepository) *ProjectService {
	return &ProjectService{repo}
}

// CreateProject creates a new project.
func (s *ProjectService) CreateProject(ctx context.Context, project *model.Project) (int, error) {
	return s.repo.CreateProject(ctx, project)
}

// GetProjectByID retrieves a project by its ID.
func (s *ProjectService) GetProjectByID(ctx context.Context, projectID, userID int) (*model.Project, error) {
	return s.repo.GetProjectByID(ctx, projectID, userID)
}

// GetAllProjects retrieves all projects for a specific user.
func (s *ProjectService) GetAllProjects(ctx context.Context, userID int) ([]*model.Project, error) {
	return s.repo.GetAllProjects(ctx, userID)
}

// UpdateProject updates an existing project.
func (s *ProjectService) UpdateProject(ctx context.Context, project *model.Project) error {
	return s.repo.UpdateProject(ctx, project)
}

// DeleteProject deletes a project by its ID.
func (s *ProjectService) DeleteProject(ctx context.Context, projectID int) error {
	return s.repo.DeleteProject(ctx, projectID)
}
