package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func main() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

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
