package entity

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewUserDtoFromLoginInput(input *LoginInput) *UserDto {
	return &UserDto{
		Email:    input.Email,
		Password: input.Password,
	}
}
