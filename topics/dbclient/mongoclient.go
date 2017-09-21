package dbclient

import (
	"fmt"
	"log"
	"time"

	"github.com/devopsdays/gather-flag/topics/model"
	"github.com/icrowley/fake"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// IMongoClient is an interface for Mongo
type IMongoClient interface {
	OpenMongoDB()
	QueryTopic(topicID string) (model.Topic, error)
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

// Seed creates a bunch of fake topics for fun
func (mc *MongoClient) Seed() {

	total := 100
	for i := 0; i < total; i++ {
		key := bson.NewObjectId()
		acc := model.Topic{
			ID:          key,
			Title:       fake.Sentence(),
			Description: fake.Paragraph(),
			Votes:       0,
			When:        time.Now(),
		}

		// jsonBytes, _ := json.Marshal(acc)

		if err := mc.mongoDB.DB("gather-flag").C("topics").Insert(acc); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Seeded %v fake topics...\n", total)
}

func (mc *MongoClient) QueryTopic(topicID string) (model.Topic, error) {
	topic := model.Topic{}
	// objectID := fmt.Sprintf("ObjectId(\"%v\")", topic)
	objectID := bson.ObjectIdHex(topicID)
	err := mc.mongoDB.DB("gather-flag").C("topics").Find(bson.M{"_id": objectID}).One(&topic)

	if err != nil {
		return model.Topic{}, err
	}

	return topic, nil

}
