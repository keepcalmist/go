package main

import (
	"net/http"

	"github.com/asaskevich/govalidator"
)

func validate(r *http.Request) string {
	return govalidator.WhiteList(r.FormValue("request"), "a-zA-Z0-9")
}
