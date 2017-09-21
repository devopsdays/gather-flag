package dbclient

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/devopsdays/gather-flag/rooms/model"
	"github.com/icrowley/fake"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// IMongoClient is an interface for Mongo
type IMongoClient interface {
	OpenMongoDB()
	QueryRoom(roomID string) (model.Room, error)
	Seed()
}

// MongoClient defines a mongo session
type MongoClient struct {
	mongoDB *mgo.Session
}

// OpenMongoDB provides a connection to a mongoDB instance
func (mc *MongoClient) OpenMongoDB() {
	var err error
	mc.mongoDB, err = mgo.Dial("db") //TODO: Make this not hard-coded
	if err != nil {
		log.Fatal(err)
	}
}

// Seed creates a bunch of fake rooms for fun
func (mc *MongoClient) Seed() {

	total := 100
	for i := 0; i < total; i++ {
		key := bson.NewObjectId()
		acc := model.Room{
			ID:       key,
			Roomname: fake.Word(),
			Capacity: (rand.Intn(100) + 1),
			Size:     "Large",
			When:     time.Now(),
		}

		// jsonBytes, _ := json.Marshal(acc)

		if err := mc.mongoDB.DB("gather-flag").C("rooms").Insert(acc); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Seeded %v fake rooms...\n", total)
}

func (mc *MongoClient) QueryRoom(roomID string) (model.Room, error) {
	room := model.Room{}
	// objectID := fmt.Sprintf("ObjectId(\"%v\")", room)
	objectID := bson.ObjectIdHex(roomID)
	err := mc.mongoDB.DB("gather-flag").C("rooms").Find(bson.M{"_id": objectID}).One(&room)

	if err != nil {
		return model.Room{}, err
	}

	return room, nil

}
