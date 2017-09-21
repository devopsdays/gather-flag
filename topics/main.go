package main

import (
	"fmt"

	"github.com/devopsdays/gather-flag/topics/dbclient"
	"github.com/devopsdays/gather-flag/topics/service"
)

var appName = "topicservice"

func main() {

	fmt.Printf("Starting %v\n", appName)
	initializeMongoClient()
	service.StartWebServer("8082")

}

// connect to db
func initializeMongoClient() {
	service.DBClient = &dbclient.MongoClient{}
	service.DBClient.OpenMongoDB()
	service.DBClient.Seed()
}
