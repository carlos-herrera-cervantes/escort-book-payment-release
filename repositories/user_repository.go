package repositories

import (
	"context"
	"escort-book-payment-release/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	GetById(ctx context.Context, id primitive.ObjectID) models.User
}

type UserRepository struct {
	Collection string
	Db         *mongo.Database
}

func (r *UserRepository) GetById(ctx context.Context, id primitive.ObjectID) models.User {
	collection := r.Db.Collection(r.Collection)

	result := collection.FindOne(ctx, bson.M{"_id": id})
	var user models.User
	result.Decode(&user)

	return user
}
