package kbb

import (
	"fmt"
	"net/http"
)


type AdData struct {
	Keywords []string
}

func GetKBBAd(w http.ResponseWriter, r *http.Request) {
	//calling ad.html, iframes src gives me queryparams, accessing those query params
	
    q := r.URL.Query()
    publisherURL := q.Get("publisherUrl")
	tsize:=q.Get("tsize")
	publisherPageTitle:=q.Get("publisherPageTitle")

	keywordUrl,_:=GetUrl("2","2",publisherURL,publisherPageTitle,tsize)

	fmt.Println(keywordUrl)
	keywords :=GetKeywords(keywordUrl)
	
	RenderTemplate(w, r ,keywords)
	
}

