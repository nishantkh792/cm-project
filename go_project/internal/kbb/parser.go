package kbb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//parses KBB url response and store them in structs and then return keywords list.

//all structures to reach xt
//structure for response of api https://contextual-stage.media.net/test/mock/kbb/response.txt
type Root struct {
	R [][]RItem `json:"r"`
}

type RItem struct {
	P  int    `json:"p"`
	Bg []BgItem `json:"bg"`
}

type BgItem struct {
	C int     `json:"c"`
	K []KItem `json:"k"`
}

type KItem struct {
	I     int64  `json:"i"`
	T     string `json:"t"`
	
	Xt    string `json:"xt"`
	Kc    string `json:"kc"`
	Kcid  int    `json:"kcid"`
	Kmd   Kmd    `json:"kmd"`
	Desc  string `json:"desc"`
	DtId  interface{} `json:"dtId"` // can be int or null
}

type Kmd struct {
	F1 string `json:"f1"`
}




func GetKeywords(keywordUrl string) ([]string){
	
	//calling the kbmock api getting response storing response into structs accessing multiple xt
	//xt denotes keyword
	// STEP 1: Call API
	resp, err := http.Get(keywordUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}


	var root Root
	if err := json.Unmarshal(body, &root); err != nil {
		panic(err)
	}

	xts := []string{}
	for _, arr := range root.R {
		for _, rItem := range arr {
			for _, bg := range rItem.Bg {
				for _, k := range bg.K {
					xts = append(xts, k.T)
				}
			}
		}
	}

	fmt.Println("XT values:")
	for i, v := range xts {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	return xts
}