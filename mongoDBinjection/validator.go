package main

import (
	"net/http"

	"github.com/asaskevich/govalidator"
)

//валидатор
func validate(r *http.Request) string {
	return govalidator.WhiteList(r.FormValue("request"), "a-zA-Z0-9")
}

func mapToString(c []*user) string {
	var endStr string
	for _, i := range c {

	}
}
