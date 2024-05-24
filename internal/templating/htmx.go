package templating

import (
	"html/template"
	"log"
	"net/http"
)

func CheckHTMX(r *http.Request) bool {
	value := r.Header.Get("HX-Request")
	if value == "" {
		return false
	}

	if value != "true" {
		log.Println("Recieved 'HX-Request' header value that is not 'true':", value)
		return false
	}

	return true
}

func RenderHTMX(w http.ResponseWriter, s int, n string, d any, t ...string) {
	w.WriteHeader(s)

	tmpl := template.New(n)
	tmpl = AddTemplates(tmpl, t...)
	tmpl.Execute(w, d)
}

func RedirectHTMX(w http.ResponseWriter, s int, url string) {
	w.WriteHeader(s)
	w.Header().Add("HX-Redirect", url)
}
