package cmd

import (
	"ad_serving/app/pkg/yahoo"
)

func Yahooflow(yahooinput *yahoo.Yahooinput) ([]map[string]string, error) {
	yahoo := yahoo.NewYahoomanager(yahooinput)
	yahoodatalist, _ := yahoo.GetYahooDataList()

	return yahoodatalist, nil

}
