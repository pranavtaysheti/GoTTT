package templating

import (
	"html/template"
	"io"
	"log"
)

const websiteTitle = "PT_XO"
const templatesPath = "templates/"

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
).ParseFiles(templatesPath + "layout.html"))

func getLayoutTmpl() *template.Template {
	t, err := layoutTmpl.Clone()
	if err != nil {
		log.Fatalln(err)
	}

	return t
}

func getTemplate(tn ...string) *template.Template {
	for i, n := range tn {
		tn[i] = templatesPath + n
	}

	tmpl := getLayoutTmpl()
	return template.Must(tmpl.ParseFiles(tn...))
}

func ExecuteLayout(w io.Writer, c any, tn ...string) error {
	t := getTemplate(tn...)

	return t.Execute(w, Layout{
		WebPages:     map[string]string{},
		Title:        "Login",
		Content:      c,
		CurrentRoute: "/",
	})
}
