package repositories

import (
	"context"
	"escort-book-payment-release/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IServiceRepository interface {
	GetById(ctx context.Context, id string) models.Service
	UpdateById(ctx context.Context, id string, service *models.Service)
}

type ServiceRepository struct {
	Collection string
	Db         *mongo.Database
}

func (r *ServiceRepository) GetById(ctx context.Context, id string) models.Service {
	parsed, _ := primitive.ObjectIDFromHex(id)
	collection := r.Db.Collection(r.Collection)

	lookupStage := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: "cards"},
			{Key: "localField", Value: "cardId"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "card"},
		},
	}}
	matchStage := bson.D{{
		Key:   "$match",
		Value: bson.D{{Key: "_id", Value: parsed}},
	}}

	cursor, _ := collection.Aggregate(ctx, mongo.Pipeline{lookupStage, matchStage})
	var services []models.Service
	cursor.All(ctx, &services)

	if len(services) == 0 {
		return models.Service{}
	}

	return services[0]
}

func (r *ServiceRepository) UpdateById(ctx context.Context, id string, document interface{}) {
	parsed, _ := primitive.ObjectIDFromHex(id)
	collection := r.Db.Collection(r.Collection)

	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	collection.FindOneAndUpdate(ctx, bson.M{"_id": parsed}, bson.M{"$set": document}, &opt)
}
