package models

type User struct {
	Id            string `bson:"_id"`
	FirebaseToken string `bson:"firebaseToken"`
	Email         string `bson:"email"`
}
