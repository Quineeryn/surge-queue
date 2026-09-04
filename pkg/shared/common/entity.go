package common

import "time"

type CommonEntity struct {
	ID string

	CreatedAt time.Time
	UpdatedAt time.Time
}
