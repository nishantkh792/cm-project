package controllers

import (
	"ad_serving/app/cmd"
	"ad_serving/app/pkg/kbb"
	"ad_serving/app/pkg/yahoo"
	"fmt"
	"html/template"
	"net/http"
)

type AdData struct {
	Keywords []string
}
type YahooData struct {
	YahooDataList []map[string]string
}

func KBBControllers(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	publisherURL := q.Get("publisherUrl")
	tsize := q.Get("tsize")
	publisherPageTitle := q.Get("publisherPageTitle")
	keywordlist := cmd.Keywordflow(&kbb.KBBinput{
		PublisherURL: publisherURL,
		Actno:        "5",
		Maxno:        "5",
		Rurl:         publisherURL,
		Ptitle:       publisherPageTitle,
		Tsize:        tsize,
	})
	tmpl := template.Must(template.ParseFiles("app/pkg/static/ad.html"))
	data := AdData{
		Keywords: keywordlist,
	}

	tmpl.Execute(w, data)

}

func YahooControllers(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	keywords := q.Get("keywords")
	// tsize := q.Get("tsize")
	// publisherPageTitle := q.Get("publisherPageTitle")
	yahoodatalist, _ := cmd.Yahooflow(&yahoo.Yahooinput{
		Keywords: keywords,
	})
	fmt.Println(yahoodatalist, "getting from ad.html")
	fmt.Println(yahoodatalist[0]["ClickUrl"])

	tmpl := template.Must(template.ParseFiles("app/pkg/static/serp.html"))
	data := YahooData{
		YahooDataList: yahoodatalist,
	}

	tmpl.Execute(w, data)

}
