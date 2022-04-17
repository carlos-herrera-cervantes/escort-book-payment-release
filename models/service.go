package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	CardId     primitive.ObjectID `bson:"cardId,omitempty"`
	CustomerId primitive.ObjectID `bson:"customerId,omitempty"`
	EscortId   primitive.ObjectID `bson:"escortId,omitempty"`
	Price      int32              `bson:"price,omitempty"`
	Status     string             `bson:"status,omitempty"`
	Card       []Card             `bson:"card,omitempty"`
}
