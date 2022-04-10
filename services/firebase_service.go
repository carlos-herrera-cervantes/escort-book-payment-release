package services

import (
	"context"
	"log"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

type IFirebaseService interface {
	SendToDevice(ctx context.Context, token, text string)
}

type FirebaseService struct{}

var client *messaging.Client
var lock = &sync.Mutex{}

func getFirebaseApp() *messaging.Client {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()

		if client == nil {
			app, _ := firebase.NewApp(context.TODO(), nil)
			client, _ = app.Messaging(context.TODO())
		}
	}

	return client
}

func (f *FirebaseService) SendToDevice(ctx context.Context, token, text string) {
	client := getFirebaseApp()

	message := &messaging.Message{
		Data: map[string]string{
			"text": text,
		},
		Token: token,
	}

	response, _ := client.Send(ctx, message)

	log.Println(response)
}
