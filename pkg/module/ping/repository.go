package ping

import (
	"context"
	"myAPI/pkg/entity"
	"sync"
	"time"
)

type PingRepository interface {
	FetchMessage(ctx context.Context) ([]entity.PingDto, error)
	CreateMessage(req *entity.PingDto) *entity.PingDto
}

type pingRepository struct {
	mu     sync.RWMutex
	data   map[int]entity.PingDto
	nextID int
}

func NewPingRepository() PingRepository {
	return &pingRepository{
		data:   make(map[int]entity.PingDto),
		nextID: 1,
	}
}

func (p *pingRepository) FetchMessage(ctx context.Context) ([]entity.PingDto, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	list := make([]entity.PingDto, 0, len(p.data))
	for _, val := range p.data {
		list = append(list, val)
	}

	select {
	case <-time.After(3 * time.Second):
		return list, nil
	case <-ctx.Done():
		return nil, ctx.Err()

	}

}

func (p *pingRepository) CreateMessage(req *entity.PingDto) *entity.PingDto {

	p.mu.Lock()

	p.data[p.nextID] = *req
	p.nextID++

	p.mu.Unlock()

	return req
}
