package entity

type PingDto struct {
	Message string
}

type PingInput struct {
	Message string `json:"message"`
}

func NewPingDtoFromInput(input *PingInput) *PingDto {
	return &PingDto{
		Message: input.Message,
	}
}
