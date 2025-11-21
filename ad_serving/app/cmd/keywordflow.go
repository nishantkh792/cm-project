package cmd

import (
	"ad_serving/app/pkg/kbb"
	"ad_serving/app/pkg/template"
	"fmt"
)

func Keywordflow(PublisherURL string, Actno string, Maxno string, Rurl string, Ptitle string, Tsize string) []string {
	kbb := kbb.NewKBBmanager(PublisherURL, Actno, Maxno, Rurl, Ptitle, Tsize)
	keywordlist := kbb.GetKeywordList()

	templateManager := template.NewTemplatemanager()
	htmltemplate, _ := templateManager.LoadTemplate()
	fmt.Println(htmltemplate)
	return keywordlist

}
