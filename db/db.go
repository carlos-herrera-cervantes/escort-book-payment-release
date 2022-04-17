package db

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var mongoOnce sync.Once

func Connect(db string) *mongo.Database {
	mongoOnce.Do(func() {
		client, err := mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(os.Getenv("MONGODB_HOST")),
		)

		if err != nil {
			log.Fatal(err)
		}

		clientInstance = client
	})

	log.Println("Successfully conected to MongoDB")

	return clientInstance.Database(db)
}
