package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var FcmClient *messaging.Client

func FirebaseInit(ctx context.Context) {
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	// Create a new firebase app
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Errorf("Error in initializing new firebase app - %v\n", err)
	}
	// Get the FCM object
	FcmClient, err = app.Messaging(ctx)
	if err != nil {
		fmt.Errorf("Error in creating firebase messaging client - %v\n", err)
	}
}
