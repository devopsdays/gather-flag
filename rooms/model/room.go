package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Room struct {
	ID       bson.ObjectId `json:"accountid" bson:"_id"`
	Roomname string        `json:"roomname" bson:"roomname"`
	Capacity int           `json:"capacity" bson:"capacity"`
	Size     string        `json:"size" bson:"size"`
	When     time.Time     `json:"when" bson:"when"`
}
