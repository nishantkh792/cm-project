package controllers

import (
	"ad_serving/app/cmd"
	"html/template"
	"net/http"
)

type AdData struct {
	Keywords []string
}

func Kbbcontrollers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	publisherURL := q.Get("publisherUrl")
	tsize := q.Get("tsize")
	publisherPageTitle := q.Get("publisherPageTitle")

	keywordlist := cmd.Keywordflow(publisherURL, "5", "5", publisherURL, publisherPageTitle, tsize)
	tmpl := template.Must(template.ParseFiles("app/pkg/static/ad.html"))
	data := AdData{
		Keywords: keywordlist,
	}

	tmpl.Execute(w, data)

}
