package activity

import (
	"context"
	"myAPI/pkg/entity"
	"myAPI/pkg/entity/query"
)

type Service interface {
	Create(ctx context.Context, log *entity.ActivityLogDto) error
	GetTransferSummary(ctx context.Context, req *query.TransferSummary) (*query.TransferSummary, error)
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

func (s *service) GetTransferSummary(ctx context.Context, req *query.TransferSummary) (*query.TransferSummary, error) {
	return s.repo.GetTransferSummary(ctx, req.UserID)
}
