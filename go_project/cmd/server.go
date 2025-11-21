package main

import (
	"fmt"
	"net/http"
	"go_project/internal/kbb"
)


//my own function that takes filename and serves the file now files all in one folder
func serveFile(w http.ResponseWriter, r *http.Request,filename string) {
	
		http.ServeFile(w, r, filename)
}

func main() {
	//routes
	
	// Serve tag.js
	http.HandleFunc("/server/tag.js", func(w http.ResponseWriter, r *http.Request) {
        serveFile(w, r, "static/tag1.js")
    })

	// Serve ad.html
	http.HandleFunc("/server/ad.html", kbb.GetKBBAd)



	// Serve ad.js
	http.HandleFunc("/server/ad.js", func(w http.ResponseWriter, r *http.Request) {
		 serveFile(w, r, "ad.js")
	})

	fmt.Println("🚀 Go Ad Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}









// //gets publisher page domain from pageurl
// func getDomain(publisherURL string) (string, error) {
//     u, err := url.Parse(publisherURL)
//     if err != nil {
//         return "", err
//     }

//     return u.Host, nil // e.g. "www.forbes.com"
// }


// //taking keyword url and modify its some of imp keys queryparams and keep else as it is
// func modifyParams(rawURL string,rurl string,d string,ptitle string,tsize string) (string, error) {
//     u, err := url.Parse(rawURL)
//     if err != nil {
//         return "", err
//     }

//     q := u.Query()





//     // Update values
//     updateKeys := map[string]string{
// 		"actno":"5",
// 		"cc":"US",
// 		"d":d,
// 		"maxno":"5",
// 		"ptitle":ptitle,
// 		"rurl":rurl,
// 		"tsize":tsize,
//     }

//     for k, v := range updateKeys {
//         q.Set(k, v)
//     }

//     u.RawQuery = q.Encode()
//     return u.String(), nil
// }


// //all structures to reach xt
// //structure for response of api https://contextual-stage.media.net/test/mock/kbb/response.txt
// type Root struct {
// 	R [][]RItem `json:"r"`
// }

// type RItem struct {
// 	P  int    `json:"p"`
// 	Bg []BgItem `json:"bg"`
// }

// type BgItem struct {
// 	C int     `json:"c"`
// 	K []KItem `json:"k"`
// }

// type KItem struct {
// 	I     int64  `json:"i"`
// 	T     string `json:"t"`
	
// 	Xt    string `json:"xt"`
// 	Kc    string `json:"kc"`
// 	Kcid  int    `json:"kcid"`
// 	Kmd   Kmd    `json:"kmd"`
// 	Desc  string `json:"desc"`
// 	DtId  interface{} `json:"dtId"` // can be int or null
// }

// type Kmd struct {
// 	F1 string `json:"f1"`
// }


// func getKeywords(keywordUrl string) ([]string){
	
// 	//calling the kbmock api getting response storing response into structs accessing multiple xt
// 	//xt denotes keyword
// 	// STEP 1: Call API
// 	resp, err := http.Get(keywordUrl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}


// 	var root Root
// 	if err := json.Unmarshal(body, &root); err != nil {
// 		panic(err)
// 	}

// 	xts := []string{}
// 	for _, arr := range root.R {
// 		for _, rItem := range arr {
// 			for _, bg := range rItem.Bg {
// 				for _, k := range bg.K {
// 					xts = append(xts, k.T)
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println("XT values:")
// 	for i, v := range xts {
// 		fmt.Printf("%d. %s\n", i+1, v)
// 	}
// 	return xts
// }

// type AdData struct {
// 	Keywords []string
// }
// func getAd(w http.ResponseWriter, r *http.Request) {
// 	//calling ad.html, iframes src gives me queryparams, accessing those query params
	
//     q := r.URL.Query()

  
//     publisherURL := q.Get("publisherUrl")
// 	tsize:=q.Get("tsize")
// 	publisherPageTitle:=q.Get("publisherPageTitle")
	





// 	publisherDomain,_ := getDomain(publisherURL) 
	


// 	keywordUrl,_:=modifyParams("http://g-usw1b-kwd-api-realapi.srv.media.net/kbb/keyword_api.php?actno=5&calling_source=cm&cc=US&combineExpired=1&crid=849176236&csid=8CUJM46V5&d=www.forbes.com&dtld=com&fm_skc=1&fpid=800015395&hs=3&https=1&json=1&kf=0&kwrd=0&kwrf=https://www.google.com&lid=224&lmsc=1&maxno=5&mtags=%7Bperform%2CBT1_sp%7D%2C%7Bsem%2Capp%2Cdmsedo%2Cmva%2Cstm%2Cconndigi%2Cpdeal%2Caudext%2Cconn%2Cginsu%7D&partnerid=7PRFT79UO&pid=8POJDA6W3&pstag=skenzo_test&pt=60&ptitle=%20Apple%20Analyst%20Reveals%20iPhone%2014%20Price%20Hikes&py=1&rurl=https://www.forbes.com/sites/gordonkelly/2022/07/17/apple-iphone-14-pro-max-pricing-price-increase-all-models&stag_tq_block=1&stags=skenzo_test&tsize=300x250&type=1&uftr=0&ugd=4&ykf=1",publisherURL,publisherDomain,publisherPageTitle,tsize)
// 	fmt.Println("final url",keywordUrl)

// 	keywords :=getKeywords(keywordUrl)
// 	fmt.Println(keywords,"in main")

// 	tmpl := template.Must(template.ParseFiles("ad.html"))
// 	data := AdData{
// 			Keywords: keywords,
// 		}

// 	tmpl.Execute(w, data)





// }
