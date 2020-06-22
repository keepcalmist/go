package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func switchUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err, "something wrong with db")
			return
		}
		db.Create(&user)

	case "GET":
		var user User
		vars := mux.Vars(r)
		id := vars["id"]
		db.Where("id = ?", id).Find(&user)
		fmt.Println(user, "          ", id)
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
		w.Write(j)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}
