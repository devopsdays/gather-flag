package main

import (
	"fmt"

	"github.com/devopsdays/gather-flag/rooms/dbclient"
	"github.com/devopsdays/gather-flag/rooms/service"
)

var appName = "roomservice"

func main() {

	fmt.Printf("Starting %v\n", appName)
	initializeMongoClient()
	service.StartWebServer("8081")

}

// connect to db
func initializeMongoClient() {
	service.DBClient = &dbclient.MongoClient{}
	service.DBClient.OpenMongoDB()
	service.DBClient.Seed()
}
