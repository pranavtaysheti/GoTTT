package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	WebsiteTitle string
	WebPages     map[string]string
	Title        string
	Content      any
	CurrentRoute string
}

type Login struct {
	ErrorMessage string
}

func getLayoutTmpl() *template.Template {
	t, err := layoutTmpl.Clone()
	if err != nil {
		log.Fatalln(err)
	}

	return t
}
func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := getLayoutTmpl().ParseFiles("templates/login.html")
		if err != nil {
			log.Fatalln(err)
		}

		err = t.Execute(w, Layout{
			WebsiteTitle: websiteTitle,
			WebPages:     map[string]string{},
			Title:        "Login",
			Content: Login{
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

	http.ListenAndServe(":3000", r)
}
