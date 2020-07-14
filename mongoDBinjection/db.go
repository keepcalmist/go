package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mgo "gopkg.in/mgo.v2/bson"
)

func getUsers(r *http.Request) string {
	log.Println("r['request'] = ", r.Form["request"])
	fmt.Println("validate data:", validate(r))
	request := r.FormValue("request")

	log.Println(request, len(request))
	if len(request) == 0 {
		return request
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Context has been created")
	ClientNONGO, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client has been created")
	err = ClientNONGO.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var val bson.M
	var jsonmarshl interface{}
	//query := `{"id":1}`
	err = mgo.UnmarshalJSON([]byte(fmt.Sprint(`{"id": `, request, ` }`)), &jsonmarshl)
	if err != nil {
		log.Println(err)
	}
	err = bson.Unmarshal([]byte(fmt.Sprint(`{"id": `, request, ` }`)), &val)
	if err != nil {
		log.Println(err)
	}
	localDB := ClientNONGO.Database("mydb")
	tasks := localDB.Collection("users")
	fmt.Println("bson filter: ", val, "   Should be: ", fmt.Sprint(`{"id": `, request, `}`))
	cur, err := tasks.Find(ctx, jsonmarshl) // Вместо второго аргумента получается инъекция bson.M{`$ne`: ""}
	if err != nil {
		log.Println("Injection hasnt been execute, filter:", val)
		log.Println(err)
		return ""
	}

	fmt.Println("vse ok!!! posle zaprosa")
	var results []*user

	var episode []bson.M
	if err = cur.All(ctx, &episode); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episode)
	for cur.Next(ctx) {
		var elem user
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
		}
		results = append(results, &elem)
	}
	fmt.Println(results)
	cur.Close(ctx)
	fmt.Println("vse ok!!! posle poiska")
	fmt.Println(results)
	return request

}
