package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Topic struct {
	ID          bson.ObjectId `json:"topicID" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Votes       int           `json:"votes" bson:"votes"`
	When        time.Time     `json:"when" bson:"when"`
}
