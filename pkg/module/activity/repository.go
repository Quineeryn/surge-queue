package activity

import (
	"context"
	"myAPI/pkg/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, log *entity.ActivityLogDto) error
}

type repository struct {
	coll *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		coll: db.Collection("activity_logs"),
	}
}

func (r *repository) Create(ctx context.Context, log *entity.ActivityLogDto) error {
	model := entity.NewActivityModelFromDto(log)

	_, err := r.coll.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
