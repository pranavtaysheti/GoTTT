package main

import (
	"net/http"
	"log"
)

type LoginPage struct {
	ErrorMessage string
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := getLayoutTmpl().ParseFiles("templates/login.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = ExecuteLayout(t, w, LoginPage{
		ErrorMessage: "Some Error happened!!!",
	})
	if err != nil {
		log.Println(err)
	}
}
