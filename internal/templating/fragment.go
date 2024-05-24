package templating

import (
	"errors"
	"html/template"
	"net/http"
)

type Fragment struct {
	name string
	data any
	tmpl *template.Template
}

func NewFragment(n string) *Fragment {
	return &Fragment{
		name: n,
		tmpl: template.New(n),
	}
}

func (f *Fragment) AddTemplates(tn ...string) {
	AddTemplates(f.tmpl, tn...)
}

func (f *Fragment) Render(w http.ResponseWriter, s int, d any) error {
	w.WriteHeader(s)

	if f.tmpl == nil {
		return errors.New("tried to render fragment with no template")
	}

	return f.tmpl.Execute(w, d)
}
