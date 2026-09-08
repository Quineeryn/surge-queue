package activity

import (
	"context"
	"encoding/json"
	"fmt"
	"myAPI/pkg/entity"
	"myAPI/pkg/entity/query"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	Create(ctx context.Context, log *entity.ActivityLogDto) error
	GetTransferSummary(ctx context.Context, req *query.TransferSummary) (*query.TransferSummary, error)
	InvalidateTransferSummaryCache(ctx context.Context, userID string) error
}

type service struct {
	repo  Repository
	redis *redis.Client
}

func NewService(repo Repository, redis *redis.Client) Service {
	return &service{
		repo:  repo,
		redis: redis,
	}
}

func (s *service) Create(ctx context.Context, req *entity.ActivityLogDto) error {
	return s.repo.Create(ctx, req)
}

func (s *service) GetTransferSummary(ctx context.Context, req *query.TransferSummary) (*query.TransferSummary, error) {

	cacheKey := fmt.Sprintf("transfer_summary:%s", req.UserID)

	val, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var summary query.TransferSummary
		if json.Unmarshal([]byte(val), &summary) == nil {
			return &summary, nil
		}
	}

	summary, err := s.repo.GetTransferSummary(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	if dataBytes, err := json.Marshal(summary); err == nil {
		_ = s.redis.Set(ctx, cacheKey, dataBytes, 5*time.Minute).Err()
	}

	return summary, nil
}

func (s *service) InvalidateTransferSummaryCache(ctx context.Context, userID string) error {
	cachedKey := fmt.Sprintf("transfer_summary:%s", userID)
	return s.redis.Del(ctx, cachedKey).Err()
}
