package repositories

import (
	"context"
	"escort-book-payment-release/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type IPaymentRepository interface {
	Create(ctx context.Context, payment *models.Payment) error
}

type PaymentRepository struct {
	Collection string
	Db         *mongo.Database
}

func (r *PaymentRepository) Create(ctx context.Context, payment *models.Payment) error {
	collection := r.Db.Collection(r.Collection)
	_, err := collection.InsertOne(ctx, payment)

	if err != nil {
		return err
	}

	return nil
}
