package service

import (
	"fmt"

	"github.com/fakelozic/tasker/internal/lib/aws"
	"github.com/fakelozic/tasker/internal/lib/job"
	"github.com/fakelozic/tasker/internal/repository"
	"github.com/fakelozic/tasker/internal/server"
)

type Services struct {
	Auth     *AuthService
	Job      *job.JobService
	Todo     *TodoService
	Comment  *CommentService
	Category *CategoryService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	s.Job.SetAuthService(authService)

	awsClient, err := aws.NewAWS(s)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS client: %w", err)
	}

	return &Services{
		Job:      s.Job,
		Auth:     authService,
		Todo:     NewTodoService(s, repos.Todo, repos.Category, awsClient),
		Comment:  NewCommentService(s, repos.Comment, repos.Todo),
		Category: NewCategoryService(s, repos.Category),
	}, nil
}
