package kbb

import(
	"net/http"
	"html/template"
)


//takes keywords and put them into ad.html
func RenderTemplate(w http.ResponseWriter, r *http.Request,keywords []string){

	tmpl := template.Must(template.ParseFiles("static/ad.html"))
	data := AdData{
			Keywords: keywords,
		}
	
	tmpl.Execute(w, data)
}