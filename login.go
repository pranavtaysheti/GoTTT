package main

import (
	"errors"
	"log"
	"net/http"
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

func checkPlayerCookie(r *http.Request) (*Player, error) {
	playerCookie, err := r.Cookie("ptxo_player")
	if err != nil {
		log.Println("Error reading player cookie, no player cookie found")
		return &Player{}, err
	}

	player, err := getPlayerbyUUID(playerCookie.Value)
	return player, err
}

func makePlayerCookie(p *Player) *http.Cookie {
	cookie := http.Cookie{
		Name:  "ptxo_player",
		Value: p.String(),
		Path:  "/",
	}

	return &cookie
}

func loginApiHandler(w http.ResponseWriter, r *http.Request) {
	playerName, roomName, err := validateLoginInfo(r)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	p := NewPlayer(playerName)
	http.SetCookie(w, makePlayerCookie(p))
	http.Redirect(w, r, "/r/"+roomName, http.StatusFound)
}
