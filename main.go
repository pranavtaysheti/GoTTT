package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/net/websocket"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", IndexPageHandler)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/html/notfound.html")
	})

	r.Route("/r/{room}", func(r chi.Router) {
		r.Use(LoginCookieMiddleware, RoomExistsMiddleware, RoomPermissionMiddleware)
		r.Get("/", RoomPageHandler)
	})

	r.Post("/api/auth", loginApiHandler)
	r.Post("/api/room", roomApiHandler)
	r.Handle("/ws", websocket.Handler(wsHandler))
	http.ListenAndServe(":3000", r)
}
