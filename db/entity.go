package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID               string `json:"_id"`
	ArrivalAirpot    int    `json:"arrival_airpot"`
	Carrier          string `json:"carrier"`
	DepartureAirport int    `json:"departure_airport"`
	Gate             int    `json:"gate"`
	Iataid           int    `json:"iataid"`
	Status           string `json:"status"`
	Termianl         string `json:"termianl"`
}

type NotificationToken struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserId    string             `bson:"userId" json:"userId"`
	DeviceId  string             `bson:"deviceId" json:"deviceId"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}

type Message struct {
	UserId string
	Text   string
}
