package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pranavtaysheti/GoTTT/internal/server"
	"github.com/pranavtaysheti/GoTTT/internal/templating"
)

type RoomPage struct {
	RoomName    string
	Player      string
}

func RoomExistsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("passing RoomExistsMiddleware")
		room_name := chi.URLParam(r, "room")
		room, ok := server.Rooms[room_name]
		if !ok {
			w.WriteHeader(http.StatusNoContent)
			templating.Render(w, "Room Not Found", nil, "roomnotfound.html")
			return
		}

		ctx := context.WithValue(r.Context(), "room", room)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RoomPermissionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("passing RoomPermissionMiddleware")
		con_cl := r.Context().Value("client")
		cl, ok := con_cl.(*server.Client)
		if !ok {
			http.Error(w, "Unable to extract client in RoomPermissionMiddleware", http.StatusInternalServerError)
			return
		}

		ro := r.Context().Value("room")
		con_room, ok := ro.(*server.Room)
		if !ok {
			http.Error(w, "Unable to extract room", http.StatusInternalServerError)
			return
		}

		cl_room, _ := cl.GetRoom()
		if cl_room != con_room {
			w.WriteHeader(http.StatusForbidden)
			templating.Render(w, "Not Permitted", nil, "notpermitted.html")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RoomPageHandler(w http.ResponseWriter, r *http.Request) {
	con_cl := r.Context().Value("client")
	cl, ok := con_cl.(*server.Client)
	if !ok {
		http.Error(w, "Unable to extract client in handler", http.StatusInternalServerError)
		return
	}

	cl_player, _ := cl.GetPlayer()
	templating.Render(
		w,
		"roomname here",
		RoomPage{
			Player:      cl_player.GetName(),
			RoomName:    "something",
		},
		"room.html",
	)
}
