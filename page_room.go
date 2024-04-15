package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pranavtaysheti/goTTT/templating"
)

type RoomPage struct {
	RoomName    string
	Player      string
}

func RoomExistsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("passing RoomExistsMiddleware")
		room_name := chi.URLParam(r, "room")
		room, ok := rooms[room_name]
		if !ok {
			w.WriteHeader(http.StatusNoContent)
			templating.ExecuteLayout(w, nil, "roomnotfound.html")
			return
		}

		ctx := context.WithValue(r.Context(), "room", room)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RoomPermissionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("passing RoomPermissionMiddleware")
		_c := r.Context().Value("client")
		cl, ok := _c.(*client)
		if !ok {
			http.Error(w, "Unable to extract client in RoomPermissionMiddleware", http.StatusInternalServerError)
			return
		}

		ro := r.Context().Value("room")
		room, ok := ro.(*Room)
		if !ok {
			http.Error(w, "Unable to extract room", http.StatusInternalServerError)
			return
		}

		if cl.room != room {
			w.WriteHeader(http.StatusForbidden)
			templating.ExecuteLayout(w, nil, "notpermitted.html")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RoomPageHandler(w http.ResponseWriter, r *http.Request) {
	_c := r.Context().Value("client")
	cl, ok := _c.(*client)
	if !ok {
		http.Error(w, "Unable to extract client in handler", http.StatusInternalServerError)
		return
	}

	templating.ExecuteLayout(
		w,
		RoomPage{
			Player:      cl.player.name,
			RoomName:    "something",
		},
		"room.html",
	)
}
