package main

import (
	"escort-book-payment-release/db"
	"escort-book-payment-release/handlers"
	"escort-book-payment-release/repositories"
	"escort-book-payment-release/services"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	handler := &handlers.PaymentHandler{
		ServiceRepository: &repositories.ServiceRepository{
			Collection: os.Getenv("SERVICE_COLLECTION"),
			Db:         db.Connect(os.Getenv("DEFAULT_DB")),
		},
		UserRepository: &repositories.UserRepository{
			Collection: os.Getenv("USER_COLLECTION"),
			Db:         db.Connect(os.Getenv("AUTHORIZER_DB")),
		},
		PaymentRepository: &repositories.PaymentRepository{
			Collection: os.Getenv("PAYMENT_COLLECTION"),
			Db:         db.Connect(os.Getenv("DEFAULT_DB")),
		},
		FirebaseService:    &services.FirebaseService{},
		EventBridgeService: &services.EventBridgeService{},
	}

	lambda.Start(handler.HandleRequest)
}
