// Package model provides the data model of the Topic object

package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Topic is a struct defining an Open Space topic
type Topic struct {
	ID          bson.ObjectId `json:"topicID" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Votes       int           `json:"votes" bson:"votes"`
	When        time.Time     `json:"when" bson:"when"`
}
