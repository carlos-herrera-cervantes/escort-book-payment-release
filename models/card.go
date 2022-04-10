package models

type Card struct {
	Id    string `bson:"_id,omitempty"`
	Token string `bson:"token,omitempty"`
}
