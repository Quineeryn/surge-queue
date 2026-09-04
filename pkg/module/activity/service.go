package activity

import (
	"context"
	"myAPI/pkg/entity"
)

type Service interface {
	Create(ctx context.Context, log *entity.ActivityLogDto) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, req *entity.ActivityLogDto) error {
	return s.repo.Create(ctx, req)
}
