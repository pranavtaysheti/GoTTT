package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
)

const ClientCookieName = "client_id"

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

func getClientCookie(r *http.Request) (*http.Cookie, error) {
	clientCookie, err := r.Cookie(ClientCookieName)
	if err != nil {
		log.Println("Error reading player cookie, no client cookie found")
		return nil, err
	}

	return clientCookie, nil
}

func getClientFromCookie(c *http.Cookie) (*client, error) {
	cl, err := getClientByUUIDString(c.Value)
	if err != nil {
		log.Println("Client not found, corresponding to cookie")
		return nil, err
	}
	return cl, err
}

func makeClientCookie(c uuid.UUID) *http.Cookie {
	cookie := http.Cookie{
		Name:  ClientCookieName,
		Value: c.String(),
		Path:  "/",
	}

	return &cookie
}


func loginApiHandler(w http.ResponseWriter, r *http.Request) {
	playerName, roomName, err := validateLoginInfo(r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	cu := NewClient()
	c := connections[cu]

	p := NewPlayer(playerName)
	ro := getOrMakeRoom(roomName)
	err = c.Login(p, ro)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.SetCookie(w, makeClientCookie(cu))
	http.Redirect(w, r, "/r/"+roomName, http.StatusFound)
}
