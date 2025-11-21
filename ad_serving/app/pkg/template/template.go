package template

import (
	"html/template"
)

type Templateinput struct {
}

func (templateinput *Templateinput) LoadTemplate() (*template.Template, error) {
	tmpl, err := template.ParseFiles("app/pkg/static/ad.html")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func NewTemplatemanager() *Templateinput {
	return &Templateinput{}
}
