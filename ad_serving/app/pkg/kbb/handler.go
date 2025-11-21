package kbb

// import (
// 	"ad_serving/app/pkg/template"
// 	"fmt"
// 	"net/http"
// )

// func GetKBBAd(w http.ResponseWriter, r *http.Request) {
// 	//calling ad.html, iframes src gives me queryparams, accessing those query params

// 	q := r.URL.Query()
// 	publisherURL := q.Get("publisherUrl")
// 	tsize := q.Get("tsize")
// 	publisherPageTitle := q.Get("publisherPageTitle")

// 	kbburl := NewKBBurl(publisherURL, "2", "2", publisherURL, publisherPageTitle, tsize)
// 	keywordUrl, _ := kbburl.GetUrl()

// 	fmt.Println(keywordUrl)
// 	keywords := GetKeywords(keywordUrl)

// 	template.RenderTemplate(w, r, keywords)

// }
