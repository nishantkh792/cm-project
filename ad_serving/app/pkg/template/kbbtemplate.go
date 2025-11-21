package template

import (
	"html/template"
	"net/http"
)

type AdData struct {
	Keywords []string
}

// takes keywords and put them into ad.html
func RenderTemplate(w http.ResponseWriter, r *http.Request, keywords []string) {

	tmpl := template.Must(template.ParseFiles("app/pkg/static/ad.html"))
	data := AdData{
		Keywords: keywords,
	}

	tmpl.Execute(w, data)
}
