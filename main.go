package main

import (
	"context"
	"test/db"
	"test/firebase"
	"test/router"
)

func main() {
	ctx := context.TODO()
	db.Init(ctx)
	defer func() {
		if err := db.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Server
	router.Init()

	// Firebase
	firebase.FirebaseInit(ctx)
}
