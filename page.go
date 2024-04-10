package main

import (
	"html/template"
	"log"
	"io"
)

const websiteTitle = "PT_XO"

type Layout struct {
	WebPages     map[string]string
	Title        string
	Content      any
	CurrentRoute string
}

var layoutTmpl = template.Must(template.New("layout.html").Funcs(
	map[string]any{
		"getPageTitle": func(title string) string {
			return title + " | " + websiteTitle
		},
	},
).ParseFiles("templates/layout.html"))

func getLayoutTmpl() *template.Template {
	t, err := layoutTmpl.Clone()
	if err != nil {
		log.Fatalln(err)
	}

	return t
}

func getTemplate(n string) *template.Template {
	tmpl := getLayoutTmpl()
	tmpl, err := tmpl.ParseFiles(n)
	if err != nil {
		log.Fatal(err)
	}

	return tmpl
}

func ExecuteLayout (t *template.Template, w io.Writer, c any) error {	
	return t.Execute(w, Layout{
		WebPages: map[string]string{},
		Title:    "Login",
		Content: c,
		CurrentRoute: "/",
	})
}
