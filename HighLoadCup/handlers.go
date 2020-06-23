package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	lg "github.com/sirupsen/logrus"
)

func switchUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//make better
	case "POST":
		var user User
		vars := mux.Vars(r)
		id := vars["id"]
		db.Where("id = ?", id).Find(&user)
		if user.ID == 0 {
			w.WriteHeader(404)
			return
		} else {
			w.WriteHeader(200)
		}
		err := json.NewDecoder(r.Body).Decode(&user)
		defer r.Body.Close()
		if err != nil {
			log.Println(err, "incorrect data in body")
			w.WriteHeader(400)
			return
		}
		db.Create(&user)

	case "GET":
		var user User
		vars := mux.Vars(r)
		id := vars["id"]
		db.Where("id = ?", id).Find(&user)
		if user.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(user)
		if err != nil {
			log.Println(err, "something wrong with db")
			return
		}
		w.WriteHeader(http.StatusOK)
		lg.Info(user)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func addEntity(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	entity := args["entity"]
	switch entity {
	//Done
	case "user":
		{
			var user User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				w.WriteHeader(400)
				lg.Warning(err, "incorrect data")
			}
			found := db.Find(&User{}, "id = ?", user.ID).RecordNotFound()
			if found == false {
				lg.Warning("id ", user.ID, "is already exist")
				w.WriteHeader(400)
				return
			} else {
				err := db.Create(&user).GetErrors()
				if err != nil {
					lg.Warning(err)
				}
				return
			}
		}
	//Done
	case "visit":
		{
			var visit Visit
			err := json.NewDecoder(r.Body).Decode(&visit)
			if err != nil {
				w.WriteHeader(400)
				lg.Warning(err, "incorrect data")
			}
			found := db.Find(&Visit{}, "id = ?", visit.ID).RecordNotFound()
			if found == false {
				lg.Warning("id ", visit.ID, "is already exist")
				w.WriteHeader(400)
				return
			} else {
				err := db.Create(&visit).GetErrors()
				if err != nil {
					lg.Warning(err)
				}
				return
			}
		}
	//Done
	case "location":
		{
			var location Location
			err := json.NewDecoder(r.Body).Decode(&location)
			if err != nil {
				w.WriteHeader(400)
				lg.Warning(err, "incorrect data")
			}
			found := db.Find(&Location{}, "id = ?", location.ID).RecordNotFound()
			if found == false {
				lg.Warning("id ", location.ID, "is already exist")
				w.WriteHeader(400)
				return
			} else {
				err := db.Create(&location).GetErrors()
				if err != nil {
					lg.Warning(err)
				}
				return
			}
		}
	}
}
