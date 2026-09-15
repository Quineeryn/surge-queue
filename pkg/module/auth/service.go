package auth

import (
	"context"
	"myAPI/pkg/entity"
	"myAPI/pkg/module/user"
	"myAPI/pkg/security"
)

type Service interface {
	Login(ctx context.Context, req *entity.LoginInput) (string, error)
}

type service struct {
	userSvc    user.Service
	jwtManager *security.JWTManager
}

func NewService(userSvc user.Service, jwtManager *security.JWTManager) Service {
	return &service{
		userSvc:    userSvc,
		jwtManager: jwtManager,
	}
}

func (s *service) Login(ctx context.Context, req *entity.LoginInput) (string, error) {
	dto := entity.NewUserDtoFromLoginInput(req)
	user, err := s.userSvc.FindByEmail(ctx, dto)

	if err != nil {
		return "", entity.ErrInvalidCredentials
	}

	err = security.VerifyPassword(req.Password, user.Password)
	if err != nil {
		return "", entity.ErrInvalidCredentials
	}

	token, err := s.jwtManager.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
