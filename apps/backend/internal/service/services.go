package service

import (
	"github.com/fakelozic/tasker/internal/lib/job"
	"github.com/fakelozic/tasker/internal/repository"
	"github.com/fakelozic/tasker/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
