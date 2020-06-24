package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	lg "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type User struct {
	ID        uint32 `json:"id,omitempty",gorm:"primary_key"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate uint64 `json:"birth_date"`
}

type Location struct {
	ID       uint32 `json:"id,omitempty",gorm:"primary_key"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance uint32 `json:"distance"`
}

type Visit struct {
	ID        uint32 `json:"id,omitempty",gorm:"primary_key"`
	Location  uint32 `json:"location"`
	User      uint32 `json:"user"`
	VisitedAt uint64 `json:"visited_at"`
	Mark      uint32 `json:"mark"`
}

type User_db struct {
	User
	gorm.Model
}

type Location_db struct {
	Location
	gorm.Model
}

type Visit_db struct {
	Visit
	gorm.Model
}

var db *gorm.DB

func main() {
	database, err := gorm.Open("postgres", "user=postgres password=toor dbname=mydb sslmode=disable")
	if err != nil {
		lg.Panic("somthing wrong with dbCon")
	} else {
		lg.Info("Connect has been created successful")
	}
	defer database.Close()
	db = database

	if db.HasTable(&User{}) {
		lg.Info("table 'user' has been added to db")
	} else {
		lg.Warning("db has table user")
	}

	router := mux.NewRouter()
	router.HandleFunc("/{entity}/{id:[0-9]+}", switchUser).Methods("GET", "POST")
	router.HandleFunc("/{entity}/new", addEntity).Methods("POST")
	router.HandleFunc("locations/{id:[0-9]+}/avg", locationAverage).Methods("GET")
	router.HandleFunc("users/{id:[0-9]+}/visits", visitsUser).Methods("GET")
	server := &http.Server{
		Addr:         ":8090",
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			lg.Panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	//block until we recieve signal(ctrl+c)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	server.Shutdown(ctx)
	lg.Info("Shutting down")
	os.Exit(0)

}
