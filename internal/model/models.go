package model

import "time"

type Event struct {
	Id     int       `bson:"_id"`
	Code   string    `bson:"code"`
	UserId string    `bson:"user_id"`
	Ts     time.Time `bson:"ts"`
}
