package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/pranavtaysheti/GoTTT/internal/server"
	"github.com/pranavtaysheti/GoTTT/internal/templating"
)


func validateLoginInfo(r *http.Request) (pn string, rn string, err error) {
	err = r.ParseForm()
	if err != nil {
		return pn, rn, err
	}

	pn = r.FormValue("player-name")
	rn = r.FormValue("room-name")

	if pn == "" {
		return pn, rn, errors.New("field player-name is empty")
	}

	if rn == "" {
		return pn, rn, errors.New("field room-name is empty")
	}

	return pn, rn, err
}


func loginApiHandler(w http.ResponseWriter, r *http.Request) {
	playerName, roomName, err := validateLoginInfo(r)
	if err != nil {
		log.Println(err)
		templating.RenderHTMX(w, http.StatusUnprocessableEntity,
			"LoginForm",
			LoginContent{
				LoginForm: LoginForm{
					Error: LoginError{
						ErrorMessage: "Error validating input data. Make sure you have not left any field blank.",
					},
					Values: LoginFormValues{
						PlayerName: playerName,
						RoomName: roomName,
					},
				},
			}, 
			"login.html")
		return
	}

	cu := server.NewClient()
	c := server.Connections[cu]

	p := server.NewPlayer(playerName)
	ro := server.GetOrMakeRoom(roomName)
	err = c.Login(p, ro)
	if err != nil {
		log.Println(err)
		//TODO: Write HTMX temlate to client
		return
	}

	http.SetCookie(w, makeClientCookie(cu))
	templating.RedirectHTMX(w, http.StatusFound, "/r/"+roomName)
}

