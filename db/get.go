package db

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetFlightDetails(flightID int) Flight {
	var f Flight
	coll := Client.Database(dbName).Collection(flightsCollection)
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{Iataid, flightID}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No flight was found with the id %s\n", flightID)
		return f
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &f)
	return f
}
