package db

// [Caution]: script to manually add flights to DB. Use only once.
/*
func AddFlightsToDB(client *mongo.Client) {
	coll := client.Database("flight_status_app").Collection("flights")
	docs := []interface{}{
		bson.D{{"iataid", 1},
			{"status", on_time},
			{"arrival_airpot", 1},
			{"departure_airport", 2},
			{"carrier", indigo_airlines},
			{"termianl", "A"},
			{"gate", 4}},
		bson.D{{"iataid", 2},
			{"status", on_time},
			{"arrival_airpot", 2},
			{"departure_airport", 3},
			{"carrier", indigo_airlines},
			{"termianl", "B"},
			{"gate", 3}},
		bson.D{{"iataid", 3},
			{"status", on_time},
			{"arrival_airpot", 3},
			{"departure_airport", 1},
			{"carrier", indigoAirlines},
			{"termianl", "A"},
			{"gate", 1}},
	}
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	list_ids := result.InsertedIDs
	fmt.Printf("Documents inserted: %v\n", len(list_ids))
	for _, id := range list_ids {
		fmt.Printf("Inserted document with _id: %v\n", id)
	}
}
*/
