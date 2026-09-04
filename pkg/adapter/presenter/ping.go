package presenter

import "myAPI/pkg/entity"

type PingOutput struct {
	Message string `json:"message"`
}

func NewPingOutput(dto *entity.PingDto) *PingOutput {
	return &PingOutput{
		Message: dto.Message,
	}
}
