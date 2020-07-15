package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const (
	mongoDB     = "MongoDB"
	description = "Simple mongoDB injection"
	myText      = `Здравствуйте, хочу представить вам nosql инъекцию 
    с использованием базы данных MongoDB и языком программирования
    Golang.`
)

type StartPage struct {
	Title       string
	Description string
	Message     string
}

type InjPage struct {
	Title       string
	Description string
	SomeText    string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	data := &StartPage{
		Title:       mongoDB,
		Description: description,
		Message:     myText,
	}
	tmpl, err := template.ParseFiles("templates/title.html")
	if err != nil {
		fmt.Fprint(w, err)
	}
	tmpl.Execute(w, data)
}

func (connect *myDBAndOthers) requestPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	data := &InjPage{
		Title:       mongoDB,
		Description: description,
		SomeText:    "lol",
	}

	tmpl, err := template.ParseFiles("templates/injection.html")
	if err != nil {
		fmt.Fprint(w, err)
	}
	tmpl.Execute(w, data)

	fmt.Println("getusers..")
	getUsers(r, connect)
	// tmpl.Execute(w, data)

}
