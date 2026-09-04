package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommonModel struct {
	ID string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *CommonModel) BeforeCreate(db *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}
