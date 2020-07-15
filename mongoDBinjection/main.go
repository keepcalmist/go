package main

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
	ID1  primitive.ObjectID `bson:"_id,omitempty"`
	ID   int                `bson:"id"`
	Name string             `bson:"name"`
	Age  uint32             `bson:"age"`
	City string             `bson:"city"`
}

type myDBAndOthers struct {
	Con *mongo.Client
}

func main() {

	var connect myDBAndOthers
	var err error
	ctx := context.Background()
	connect.Con, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = connect.Con.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/injection", connect.requestPage)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
