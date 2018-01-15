package domain

import (
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"
)

type GetUserID struct{}

func (g *GetUserID) Execute(app *firebase.App, token string) (string, error) {
	authToken, err := verifyIDToken(app, token)
	return authToken.UID, err
}

func verifyIDToken(app *firebase.App, idToken string) (*auth.Token, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	return token, err
}
