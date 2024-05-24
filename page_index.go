package main

import (
	"net/http"

	"github.com/pranavtaysheti/GoTTT/internal/templating"
)

type LoginContent struct {
	LoginForm LoginForm
}

type LoginError struct {
	ErrorMessage string
}

type LoginFormValues struct {
	PlayerName string
	RoomName string
}

type LoginForm struct {
	Values LoginFormValues
	Error LoginError
}

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {

	
	templating.Render(
		w,
		"Login",
		LoginContent{
			LoginForm: LoginForm{

			},
		},
		"login.html",
	)
}
