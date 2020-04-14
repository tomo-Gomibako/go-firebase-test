package server

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var authContext context.Context

// Run run server
func Run() {
	authContext = context.Background()

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("public", true)))
	api := router.Group("/api")
	api.GET("/ping", ping)
	router.Run()
}

func ping(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if !verifyToken(token) {
		ctx.JSON(403, &gin.H{
			"error": "invalid token",
		})
		return
	}
	ctx.String(200, "pong")
}

func verifyToken(token string) bool {
	client := getClient()
	_, err := client.VerifyIDToken(authContext, token)
	if err != nil {
		return false
	}
	return true
}

// func createUser(user *auth.UserToCreate) *auth.UserRecord {
// 	client := getClient()
// 	record, err := client.CreateUser(authContext, user)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return record
// }

func getApp() firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return *app
}

func getClient() auth.Client {
	app := getApp()
	client, err := app.Auth(authContext)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return *client
}
