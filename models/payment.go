package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	EscortId    primitive.ObjectID `bson:"escortId"`
	CustomerId  primitive.ObjectID `bson:"customerId"`
	ServiceId   primitive.ObjectID `bson:"serviceId"`
	LogRequest  string             `bson:"logRequest"`
	LogResponse string             `bson:"logResponse"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

func (p *Payment) SetDefaultValues() {
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = time.Now().UTC()
}
