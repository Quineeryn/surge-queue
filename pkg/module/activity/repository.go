package activity

import (
	"context"
	"myAPI/pkg/entity"
	"myAPI/pkg/entity/query"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, log *entity.ActivityLogDto) error
	GetTransferSummary(ctx context.Context, userID string) (*query.TransferSummary, error)
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

func (r *repository) GetTransferSummary(ctx context.Context, userID string) (*query.TransferSummary, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"user_id": userID,
			"action":  "TRANSFER",
		}}},
		{{Key: "$group", Value: bson.M{
			"_id":                "$user_id",
			"total_amount":       bson.M{"$sum": "$amount"},
			"total_transactions": bson.M{"$sum": 1},
		}}},
	}

	cursor, err := r.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []query.TransferSummary
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return &query.TransferSummary{UserID: userID, TotalAmount: 0, TotalTransactions: 0}, nil
	}
	return &result[0], nil
}
