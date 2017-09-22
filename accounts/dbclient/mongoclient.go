// Package dbclient contain the required logic for communication with the
// database. All data operations are included.
package dbclient

import (
	"fmt"
	"log"
	"time"

	"github.com/devopsdays/gather-flag/accounts/model"
	"github.com/icrowley/fake"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// IMongoClient is an interface for Mongo
type IMongoClient interface {
	OpenMongoDB()
	QueryAccount(accountID string) (model.Account, error)
	ListAccounts() ([]model.Account, error)
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

// Seed creates a bunch of fake accounts for fun
func (mc *MongoClient) Seed() {

	total := 100
	for i := 0; i < total; i++ {
		key := bson.NewObjectId()
		acc := model.Account{
			ID:        key,
			Username:  fake.UserName(),
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Email:     fake.EmailAddress(),
			When:      time.Now(),
		}

		// jsonBytes, _ := json.Marshal(acc)

		if err := mc.mongoDB.DB("gather-flag").C("accounts").Insert(acc); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}

// QueryAccount connects to the database and pulls back a single account based upon
// the accountID as provided
func (mc *MongoClient) QueryAccount(accountID string) (model.Account, error) {
	account := model.Account{}
	// objectID := fmt.Sprintf("ObjectId(\"%v\")", account)
	objectID := bson.ObjectIdHex(accountID)
	err := mc.mongoDB.DB("gather-flag").C("accounts").Find(bson.M{"_id": objectID}).One(&account)

	if err != nil {
		return model.Account{}, err
	}

	return account, nil

}

// ListAccounts connectes to the database and returns all accounts
func (mc *MongoClient) ListAccounts() ([]model.Account, error) {
	res := []model.Account{}
	if err := mc.mongoDB.DB("gather-flag").C("accounts").Find(nil).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
