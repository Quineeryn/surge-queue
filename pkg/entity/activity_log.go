package entity

import (
	"myAPI/pkg/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActivityLogDto struct {
	ID        primitive.ObjectID
	UserID    string
	Action    string
	Amount    int
	MetaData  map[string]any
	CreatedAt time.Time
}

func NewActivityLogDtoFromModel(log *model.ActivityLog) *ActivityLogDto {
	return &ActivityLogDto{
		ID:        log.ID,
		UserID:    log.UserID,
		Action:    log.Action,
		Amount:    log.Amount,
		MetaData:  log.MetaData,
		CreatedAt: log.CreatedAt,
	}
}

func NewActivityModelFromDto(dto *ActivityLogDto) *model.ActivityLog {
	return &model.ActivityLog{
		ID:        dto.ID,
		UserID:    dto.UserID,
		Action:    dto.Action,
		Amount:    dto.Amount,
		MetaData:  dto.MetaData,
		CreatedAt: dto.CreatedAt,
	}
}
