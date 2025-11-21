//taking keyword url and modify its some of imp keys queryparams and keep else as it is
//gets publisher page domain from pageurl

package kbb

import (
	"fmt"
	"net/url"
)

//query params struct then pass each query params struct to all its functions

type Params struct {
	PublisherURL string
	Actno        string
	Maxno        string
	Rurl         string
	Ptitle       string
	Tsize        string
}

func (p *Params) GetDomain() (string, error) {
	u, err := url.Parse(p.PublisherURL)
	if err != nil {
		return "", err
	}

	return u.Host, nil
}

func (p *Params) GetKBBUrl() (string, error) {

	baseURL := "http://g-usw1b-kwd-api-realapi.srv.media.net/kbb/keyword_api.php"

	d, _ := p.GetDomain()

	queryParams := url.Values{}
	queryParams.Add("actno", p.Actno)
	queryParams.Add("cc", "US")
	queryParams.Add("d", d)
	queryParams.Add("maxno", p.Maxno)
	queryParams.Add("ptitle", p.Ptitle)
	queryParams.Add("tsize", p.Tsize)

	queryParams.Add("kwrf", "https://www.google.com")
	queryParams.Add("calling_source", "cm")
	queryParams.Add("combineExpired", "1")
	queryParams.Add("crid", "849176236")
	queryParams.Add("csid", "8CUJM46V5")
	queryParams.Add("dtld", "com")
	queryParams.Add("fm_skc", "1")
	queryParams.Add("fpid", "800015395")
	queryParams.Add("hs", "3")
	queryParams.Add("https", "1")
	queryParams.Add("json", "1")
	queryParams.Add("kf", "0")
	queryParams.Add("kwrd", "0")
	queryParams.Add("lid", "224")
	queryParams.Add("lmsc", "1")
	queryParams.Add("mtags", "{perform,BT1_sp},{sem,app,dmsedo,mva,stm,conndigi,pdeal,audext,conn,ginsu}")
	queryParams.Add("partnerid", "7PRFT79UO")
	queryParams.Add("pid", "8POJDA6W3")
	queryParams.Add("pstag", "skenzo_test")
	queryParams.Add("pt", "60")
	queryParams.Add("py", "1")
	queryParams.Add("stag_tq_block", "1")
	queryParams.Add("stags", "skenzo_test")
	queryParams.Add("type", "1")
	queryParams.Add("uftr", "0")
	queryParams.Add("ugd", "4")
	queryParams.Add("ykf", "1")

	// Encode the query parameters
	encodedQuery := queryParams.Encode() // Output: pageNum=2&sortBy=name&userId=12345

	// Combine the base URL and query parameters
	finalURL := fmt.Sprintf("%s?%s", baseURL, encodedQuery)

	fmt.Println(finalURL) // Output: https://api.example.com/data?pageNum=2&sortBy=name&userId=12345
	return finalURL, nil
}

// taking keyword url and modify its some of imp keys queryparams and keep else as it is
func (p *Params) GetUrl() (string, error) {

	rawURL := "http://g-usw1b-kwd-api-realapi.srv.media.net/kbb/keyword_api.php?calling_source=cm&combineExpired=1&crid=849176236&csid=8CUJM46V5&dtld=com&fm_skc=1&fpid=800015395&hs=3&https=1&json=1&kf=0&kwrd=0&kwrf=https://www.google.com&lid=224&lmsc=1&mtags=%7Bperform%2CBT1_sp%7D%2C%7Bsem%2Capp%2Cdmsedo%2Cmva%2Cstm%2Cconndigi%2Cpdeal%2Caudext%2Cconn%2Cginsu%7D&partnerid=7PRFT79UO&pid=8POJDA6W3&pstag=skenzo_test&pt=60&py=1&stag_tq_block=1&stags=skenzo_test&type=1&uftr=0&ugd=4&ykf=1"
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	q := u.Query()

	d, _ := p.GetDomain()

	// Update values
	updateKeys := map[string]string{
		"actno":  p.Actno,
		"cc":     "US",
		"d":      d,
		"maxno":  p.Maxno,
		"ptitle": p.Ptitle,
		"rurl":   p.Rurl,
		"tsize":  p.Tsize,
	}

	for k, v := range updateKeys {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

func NewKBBurl(publisherURL string, actno string, maxno string, rurl string, ptitle string, tsize string) *Params {
	return &Params{
		PublisherURL: publisherURL,
		Actno:        actno,
		Maxno:        maxno,
		Rurl:         rurl,
		Ptitle:       ptitle,
		Tsize:        tsize,
	}
}
