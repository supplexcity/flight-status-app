package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert a token
func InsertToken(
	token NotificationToken,
	ctx context.Context,
) error {
	// Check if the token already exists
	coll := Client.Database(dbName).Collection(NotificationTokensCollection)
	filter := bson.D{{Key: "deviceId", Value: token.DeviceId}}
	res := coll.FindOne(ctx, filter)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			// If token does not exist insert it
			token.ID = primitive.NewObjectID()
			_, err := coll.InsertOne(ctx, token)
			return err
		}
		return res.Err()
	}

	// If token exists update the timestamp to now
	_, err := coll.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"timestamp": time.Now().UTC()}})
	return err
}

func GetTokens(userId string, ctx context.Context) ([]string, error) {
	var tokens []string
	coll := Client.Database(dbName).Collection(NotificationTokensCollection)
	filter := bson.D{{Key: "userId", Value: userId}}
	_, err := coll.Find(ctx, filter)
	if err != nil {
		fmt.Errorf("Error in fetching tokens for user - %s, error - %v", userId, err)
		return nil, err
	}
	return tokens, nil
}
