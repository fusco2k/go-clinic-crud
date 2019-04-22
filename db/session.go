package db

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//NewSession gets a mongo connection and returns a client pointer
func NewSession() *mongo.Client {

	//Connects to mongo server and gets a client pointer
	Client, err := mongo.Connect(nil, options.Client().ApplyURI("mongodb://192.168.122.95:27017"))
	if err != nil {
		log.Fatal(err)
	}

	//Pings the connection
	err = Client.Ping(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	return Client
}
