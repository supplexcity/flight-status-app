package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var StatusMapping = map[string]int{
	onTime:      100,
	cancelled:   101,
	delay:       102,
	gateChanged: 103,
}

func UpdateFlightStatus(flightID int, newStatus string) {
	// get present status of flight
	f := GetFlightDetails(flightID)
	if f.Status == newStatus {
		fmt.Printf("Status is same for flight id = %d!\n", flightID)
		return
	}
	filter := bson.D{{"iataid", flightID}}
	update := bson.D{{"$set", bson.D{{"status", newStatus}}}}

	coll := Client.Database(dbName).Collection(flightsCollection)
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Flight %d status updated\n", flightID)
	// TODO - notify the users
}
