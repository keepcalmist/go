package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	lg "github.com/sirupsen/logrus"
)

func switchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entity := vars["entity"]
	switch entity {
	case "users":
		{
			switch r.Method {
			//make better
			case "POST":
				var user User

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
				w.Header().Set("Content-Type", "application/json")
				return

			case "GET":
				var user User

				id := vars["id"]
				notFound := db.Where("id = ?", id).Find(&user).RecordNotFound()
				if notFound == true {
					w.WriteHeader(http.StatusNotFound)
					return
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
	case "locations":
		{
			switch r.Method {
			case "POST":
				{
					var location Location
					err := json.NewDecoder(r.Body).Decode(&location)
					if err != nil {
						w.WriteHeader(400)
						lg.Warning(err, "incorrect data")
						return
					}

					id := vars["id"]
					found := db.Find(&Location{}, "id = ?", id).RecordNotFound()
					if found == true {
						w.WriteHeader(404)
						return
					} else {
						w.Header().Set("Content-Type", "application/json")
						//todo end this
						err := db.Model(&Location{}).Where("id = ?", id).Updates(location).GetErrors()
						lg.Info("Updating ", id)
						if len(err) != 1 {
							lg.Warning(err)
						}
						return
					}
				}
			case "GET":
				{
					var location Location
					id := vars["id"]
					found := db.Where("id = ?", id).Find(&location).RecordNotFound()
					if found == true {
						w.WriteHeader(http.StatusNotFound)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					j, err := json.Marshal(location)
					if err != nil {
						log.Println(err, "something wrong with JSON")
						w.WriteHeader(http.StatusNotFound)
						return

					}
					w.WriteHeader(http.StatusOK)
					lg.Info(location)
					w.Write(j)

				}
			default:
				w.WriteHeader(http.StatusNotFound)

			}
		}
	case "visits":
		{
			switch r.Method {
			case "POST":
				{
					var visit Visit
					err := json.NewDecoder(r.Body).Decode(&visit)
					if err != nil {
						w.WriteHeader(400)
						lg.Warning(err, "incorrect data")
						return
					}

					id := vars["id"]
					found := db.Find(&Visit{}, "id = ?", id).RecordNotFound()
					if found == true {
						w.WriteHeader(404)
						return
					} else {
						w.Header().Set("Content-Type", "application/json")
						err := db.Model(&Visit{}).Where("id = ?", id).Updates(visit).GetErrors()
						lg.Info("Updating ", id)
						if len(err) != 1 {
							lg.Warning(err)
						}
						return
					}
				}
			case "GET":
				{
					var visit Visit
					id := vars["id"]
					found := db.Where("id = ?", id).Find(&visit).RecordNotFound()
					if found == true {
						w.WriteHeader(http.StatusNotFound)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					j, err := json.Marshal(visit)
					if err != nil {
						log.Println(err, "something wrong with JSON")
						w.WriteHeader(http.StatusNotFound)
						return

					}
					w.WriteHeader(http.StatusOK)
					lg.Info(visit)
					w.Write(j)
				}
			default:
				w.WriteHeader(http.StatusNotFound)
			}

		}

	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}

}

func addEntity(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	entity := args["entity"]
	switch entity {
	//Done
	case "users":
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
				if len(err) != 1 {
					lg.Warning(err)
				}
				return
			}
		}
	//Done
	case "visits":
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
	case "locations":
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

//locations/id/avg
func locationAverage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromDate, found := vars["fromDate"]
	toDate, found := vars["toDate"]
	fromAge, found := vars["fromAge"]
	toAge, found := vars["toAge"]
	gender, found := vars["gender"]

	id := vars["id"]

	var users []User
}

//users/id/visits?params...
func visitsUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromDate, found := vars["fromDate"]
	toDate, found := vars["toDate"]
	country, found := vars["country"]
	toDistance, found := vars["toDistance"]

}
