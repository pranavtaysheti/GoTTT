package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/net/websocket"
)

const websiteTitle = "PT_XO"

var layoutTmpl = template.Must(template.New("layout.html").Funcs(
	map[string]any{
		"getPageTitle": func(title string) string {
			return title + " | " + websiteTitle
		},
	},
).ParseFiles("templates/layout.html"))

type Layout struct {
	WebPages     map[string]string
	Title        string
	Content      any
	CurrentRoute string
}

type LoginPage struct {
	ErrorMessage string
}

type RoomPage struct {
	ClientName string
	RoomName   string
}

func getLayoutTmpl() *template.Template {
	t, err := layoutTmpl.Clone()
	if err != nil {
		log.Fatalln(err)
	}

	return t
}

func getClientByUUIDString(u string) (*client, error) {
	for id, c := range connections {
		if id.String() == u {
			return c, nil
		}
	}

	return &client{}, errors.New("Client with given uuid not found")
}

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := getLayoutTmpl().ParseFiles("templates/login.html")
		if err != nil {
			log.Fatalln(err)
		}

		err = t.Execute(w, Layout{
			WebPages: map[string]string{},
			Title:    "Login",
			Content: LoginPage{
				ErrorMessage: "Some Error happened!!!",
			},
			CurrentRoute: "/",
		})
		if err != nil {
			log.Println(err)
		}
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/html/notfound.html")
	})

	r.Get("/r/{room}", func(w http.ResponseWriter, r *http.Request) {
		room_name := chi.URLParam(r, "room")
		c, err := getClientCookie(r)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		tmpl := getLayoutTmpl()
		tmpl, err = tmpl.ParseFiles("templates/room.html")
		if err != nil {
			log.Fatal(err)
		}

		tmpl.Execute(w, Layout{
			WebPages: map[string]string{},
			Title:    "Room",
			Content: RoomPage{
				ClientName: c.Value,
				RoomName:   room_name,
			},
		})
	})

	r.Post("/api/login", loginApiHandler)
	r.Post("/api/board", boardApiHandler)
	r.Post("/api/player", playerApiHandler)
	r.Handle("/ws", websocket.Handler(wsHandler))
	http.ListenAndServe(":3000", r)
}
