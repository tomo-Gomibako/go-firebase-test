package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func main() {
	ctx := context.Background()
	signup(ctx)
}

func getClient(ctx context.Context) auth.Client {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return *client
}

func signup(ctx context.Context) {
	client := getClient(ctx)
	user := (&auth.UserToCreate{}).
		Email("test@example.com").
		EmailVerified(false).
		PhoneNumber("+819012345678").
		Password("password").
		DisplayName("George Tokoro").
		PhotoURL("http://www.example.com/george-t.png").
		Disabled(false)
	record, err := client.CreateUser(ctx, user)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	log.Printf("%#v", record.UserInfo)
}
