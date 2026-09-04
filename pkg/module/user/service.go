package user

import (
	"context"
	"myAPI/database"
	"myAPI/pkg/entity"
)

type Service interface {
	Create(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error)
	FindAll(ctx context.Context) ([]entity.UserDto, error)
	FindById(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error)
	FindByEmail(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error)
	Update(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error)
	Delete(ctx context.Context, req *entity.UserDto) error
	Transfer(ctx context.Context, fromID, toID string, amount int) error
}

type service struct {
	repo      Repository
	txManager database.TxManager
}

func NewService(repo Repository, txManager database.TxManager) Service {
	return &service{
		repo:      repo,
		txManager: txManager,
	}
}

func (s *service) Create(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error) {
	_, err := s.FindByEmail(ctx, req)
	if err == nil {
		return nil, entity.ErrEmailDuplicate
	}

	return s.repo.Create(ctx, req)
}

func (s *service) FindAll(ctx context.Context) ([]entity.UserDto, error) {
	return s.repo.FindAll(ctx)
}

func (s *service) FindById(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error) {
	return s.repo.FindById(ctx, req)
}

func (s *service) Update(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error) {
	return s.repo.Update(ctx, req)
}

func (s *service) Delete(ctx context.Context, req *entity.UserDto) error {
	return s.repo.Delete(ctx, req)
}

func (s *service) FindByEmail(ctx context.Context, req *entity.UserDto) (*entity.UserDto, error) {
	return s.repo.FindByEmail(ctx, req)
}

func (s *service) Transfer(ctx context.Context, fromID, toID string, amount int) error {
	if amount <= 0 || fromID == toID {
		return entity.ErrInvalidTransfer
	}

	return s.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		firstID, secondID := fromID, toID
		if fromID > toID {
			firstID, secondID = toID, fromID
		}

		firstUser, err := s.repo.FindByIdForUpdate(ctx, firstID)
		if err != nil {
			return err
		}

		secondUser, err := s.repo.FindByIdForUpdate(ctx, secondID)
		if err != nil {
			return err
		}

		var sender, receiver *entity.UserDto
		if firstUser.ID == fromID {
			sender, receiver = firstUser, secondUser
		} else {
			sender, receiver = secondUser, firstUser
		}

		if sender.Balance < amount {
			return entity.ErrBalanceInsufficient
		}

		if err := s.repo.UpdateBalance(ctx, fromID, sender.Balance-amount); err != nil {
			return err
		}

		if err := s.repo.UpdateBalance(ctx, toID, receiver.Balance+amount); err != nil {
			return err
		}

		return nil
	})
}
