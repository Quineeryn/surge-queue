package presenter

import "myAPI/pkg/entity"

type UserOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserOutputFromDto(dto *entity.UserDto) *UserOutput {
	if dto == nil {
		return nil
	}

	return &UserOutput{
		ID:    dto.ID,
		Name:  dto.Name,
		Email: dto.Email,
	}
}
