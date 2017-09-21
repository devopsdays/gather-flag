package main

import (
	"fmt"

	"github.com/mattstratton/blondie/accounts/dbclient"
	"github.com/mattstratton/blondie/accounts/service"
)

var appName = "accountservice"

func main() {

	fmt.Printf("Starting %v\n", appName)
	initializeMongoClient()
	service.StartWebServer("8080")

}

// connect to db
func initializeMongoClient() {
	service.DBClient = &dbclient.MongoClient{}
	service.DBClient.OpenMongoDB()
	service.DBClient.Seed()
}
