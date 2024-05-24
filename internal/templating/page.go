package templating

import (
	"html/template"
	"io"
	"log"
)

const websiteTitle = "PT_XO"
const templatesPath = "templates/"

type Layout struct {
	Title        string
	Content      any
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

func AddTemplates(tmpl *template.Template, tn ...string) *template.Template {

	for i, n := range tn {
		tn[i] = templatesPath + n
	}

	tmpl, err := tmpl.ParseFiles(tn...)
	if err != nil {
		log.Fatal(err)
	}

	return tmpl
}

func getTemplate(tn ...string) *template.Template {
	tmpl := getLayoutTmpl()
	return AddTemplates(tmpl, tn...)
}

func Render(w io.Writer, pt string, c any, tn ...string) error {
	t := getTemplate(tn...)

	return t.Execute(w, Layout{
		Title:        pt,
		Content:      c,
	})
}
