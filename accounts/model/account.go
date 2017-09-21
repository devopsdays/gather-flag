package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID        bson.ObjectId `json:"accountid" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	FirstName string        `json:"firstname" bson:"firstname"`
	LastName  string        `json:"lastname" bson:"lastname"`
	Email     string        `json:"email" bson:"email"`
	When      time.Time     `json:"when" bson:"when"`
}
