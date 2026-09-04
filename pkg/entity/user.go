package entity

import (
	"myAPI/pkg/model"
	"myAPI/pkg/shared/common"
)

type UserDto struct {
	common.CommonEntity
	Name    string
	Email   string
	Balance int
}

type UserInput struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type TransferInput struct {
	FromID string `json:"from_id" validate:"required"`
	ToID   string `json:"to_id" validate:"required"`
	Amount int    `json:"amount" validate:"required,gt=0"`
}

func NewUserDtoFromInput(input *UserInput) *UserDto {
	return &UserDto{
		Name:  input.Name,
		Email: input.Email,
	}
}

func NewUserModelFromDto(dto *UserDto) *model.User {
	return &model.User{
		CommonModel: common.CommonModel{
			ID: dto.ID,
		},
		Name:    dto.Name,
		Email:   dto.Email,
		Balance: dto.Balance,
	}
}

func NewUserDtoFromModel(m *model.User) *UserDto {
	return &UserDto{
		CommonEntity: common.CommonEntity{
			ID:        m.ID,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
		Name:    m.Name,
		Email:   m.Email,
		Balance: m.Balance,
	}
}
