package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActivityLog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Action    string             `bson:"action"`
	Amount    int                `bson:"amount,omitempty"`
	MetaData  map[string]any     `bson:"metadata,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
}
