package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	ID   int
	Name string
	Age  uint32
	City string
}

type myDBAndOthers struct {
	Client *mongo.Client
}

var ClientNONGO mongo.Client

func main() {

	//LocalDB = Client.Database("injection")
	// tasks := localDB.Collection("users")
	// city := `madrid`
	// User := user{ID: 12, Name: "Kirill", Age: 20, City: city}
	// result, err := tasks.InsertOne(ctx, User)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(result)
	// cur, err := tasks.Find(ctx, bson.M{"id": bson.M{`$ne`: ""}}) // Вместо второго аргумента получается инъекция
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var results []*user

	// var episode []bson.M
	// if err = cur.All(ctx, &episode); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(episode)
	// for cur.Next(ctx) {
	// 	var elem user
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	results = append(results, &elem)
	// }
	// fmt.Println(results)
	// cur.Close(ctx)

	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/injection", requestPage)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
