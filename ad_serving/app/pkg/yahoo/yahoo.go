package yahoo

import (
	"ad_serving/app/pkg/http"
	"ad_serving/app/pkg/kbb"
	"ad_serving/app/pkg/model"
	"encoding/xml"
	"net/url"
)

type Yahooinput struct {
	Keywords string
}

// taking keyword url and modify its some of imp keys queryparams and keep else as it is
func (yahooinput *Yahooinput) GetUrl() (string, error) {

	u, err := url.Parse(kbb.YahooUrl)
	if err != nil {
		return "", err
	}

	q := u.Query()

	// Update values
	updateKeys := map[string]string{
		"Keywords": yahooinput.Keywords,
	}

	for k, v := range updateKeys {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (yahooinput *Yahooinput) ParseResponse(body []byte) *model.Results {

	var results model.Results
	if err := xml.Unmarshal(body, &results); err != nil {
		panic(err)
	}

	return &results
}

func (yahooinput *Yahooinput) GetYahooData(root *model.Results) ([]map[string]string, error) {

	output := []map[string]string{}

	for _, l := range root.ResultSet.Listings {
		item := map[string]string{
			"Title":       l.Title,
			"Description": l.Description,
			"SiteHost":    l.SiteHost,
			"ClickUrl":    l.ClickUrl.URL,
		}

		output = append(output, item)
	}

	return output, nil
}

func (yahooinput *Yahooinput) GetYahooDataList() ([]map[string]string, error) {

	url, _ := yahooinput.GetUrl()

	httprequest := http.NewHttp()
	response, _ := httprequest.Get(url)
	parseStruct := yahooinput.ParseResponse(response)
	yahoodatalist, _ := yahooinput.GetYahooData(parseStruct)
	return yahoodatalist, nil
}
func NewYahoomanager(yahooinput *Yahooinput) *Yahooinput {
	return yahooinput
}
