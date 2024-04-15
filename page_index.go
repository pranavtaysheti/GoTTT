package main

import (
	"net/http"

	"github.com/pranavtaysheti/goTTT/templating"
)

type LoginPage struct {
	ErrorMessage string
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	templating.ExecuteLayout(
		w,
		LoginPage{
			ErrorMessage: "Some Error happened!!!",
		},
		"login.html",
	)
}
