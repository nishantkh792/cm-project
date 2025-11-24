package cmd

import (
	"ad_serving/app/pkg/kbb"
	"ad_serving/app/pkg/template"
	"fmt"
)

func Keywordflow(kbbinput *kbb.KBBinput) []string {
	kbb := kbb.NewKBBmanager(kbbinput)
	keywordlist := kbb.GetKeywordList()

	templateManager := template.NewTemplatemanager()
	htmltemplate, _ := templateManager.LoadTemplate()
	fmt.Println(htmltemplate)
	return keywordlist

}
