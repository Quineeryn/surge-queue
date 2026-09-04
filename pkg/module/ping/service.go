package ping

import (
	"context"
	"myAPI/pkg/entity"
	"time"
)

type PingService interface {
	GetMessage(ctx context.Context) ([]entity.PingDto, error)
	CreateMessage(req *entity.PingDto) *entity.PingDto
}

type pingService struct {
	repo PingRepository
}

func NewPingService(repo PingRepository) PingService {
	return &pingService{
		repo: repo,
	}
}

func (p *pingService) GetMessage(ctx context.Context) ([]entity.PingDto, error) {

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	return p.repo.FetchMessage(ctx)
}

func (p *pingService) CreateMessage(req *entity.PingDto) *entity.PingDto {
	return p.repo.CreateMessage(req)
}
