package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	mgo "gopkg.in/mgo.v2/bson"
)

func getUsers(r *http.Request, con *myDBAndOthers) []*user {
	log.Println("r['request'] = ", r.Form["request"])
	fmt.Println("validate data:", validate(r))
	request := r.FormValue("request")

	log.Println(request, len(request))
	if len(request) == 0 {
		return nil
	}
	var val bson.M
	var jsonmarshl interface{}
	err := mgo.UnmarshalJSON([]byte(fmt.Sprint(`{"id": `, request, ` }`)), &jsonmarshl) //реализация инъекции
	if err != nil {
		log.Println(err)
	}

	ctx := context.Background()
	localDB := con.Con.Database("mydb")
	tasks := localDB.Collection("users")
	cur, err := tasks.Find(ctx, jsonmarshl)
	if err != nil {
		log.Println("Injection hasnt been execute, filter:", val)
		log.Println(err)
		return nil
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
	return results

}
