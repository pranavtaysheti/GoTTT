package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"log"
)

type RoomPage struct {
	ClientName  string
	RoomName    string
	Player      string
	RoomPlayers []string
}


func RoomPageHandler(w http.ResponseWriter, r *http.Request) {
		room_name := chi.URLParam(r, "room")
		room, ok := rooms[room_name]
		if !ok {
			ExecuteLayout(getTemplate("templates/roomnotfound.html"), w, nil)
		}

		c, err := getClientCookie(r)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		cl, err := getClientFromCookie(c)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		tmpl := getLayoutTmpl()
		tmpl, err = tmpl.ParseFiles("templates/room.html")
		if err != nil {
			log.Fatal(err)
		}

		ExecuteLayout(tmpl, w, RoomPage{
			ClientName:  c.Value,
			Player:      cl.player.name,
			RoomPlayers: room.PlayerNames(),
			RoomName:    room_name,
		})
	}
