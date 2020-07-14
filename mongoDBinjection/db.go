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
)

func getUsers(r *http.Request) string {
	fmt.Println("Kak dolzhno BbIt':", bson.M{"id": bson.M{`$ne`: ""}})
	log.Println("r['request'] = ", r.Form["request"])
	request := r.FormValue("request")
	log.Println(request, len(request))
	if len(request) == 0 {
		return request
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("Context has been created")
	ClientNONGO, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kirill:KI990856u@cluster0.z1i9r.mongodb.net/<Cluster0>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client has been created")
	err = ClientNONGO.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = ClientNONGO.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongo db")
	localDB := ClientNONGO.Database("injection")
	tasks := localDB.Collection("users")
	cur, err := tasks.Find(ctx, `{"id":{$ne:0}}`) // Вместо второго аргумента получается инъекция bson.M{`$ne`: ""}
	fmt.Println("bson structer:", bson.M{"id": request})
	if err != nil {
		log.Println("Injection hasnt been execute, filter:", bson.M{"id": request})
		log.Println(err)
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
