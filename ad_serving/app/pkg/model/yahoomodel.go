package model

import "encoding/xml"

type Results struct {
	XMLName           xml.Name          `xml:"Results"`
	KeywordsRatingSet KeywordsRatingSet `xml:"KeywordsRatingSet"`
	SearchID          string            `xml:"SearchID"`
	ImpressionGUID    string            `xml:"ImpressionGUID"`
	TraceId           string            `xml:"TraceId"`
	RGUId             string            `xml:"RGUId"`
	ClientID          string            `xml:"ClientID"`
	ResultSet         ResultSet         `xml:"ResultSet"`
}

type KeywordsRatingSet struct {
	Keywords string `xml:"keywords,attr"`
	Market   string `xml:"Market"`
}

type ResultSet struct {
	ID          string    `xml:"id,attr"`
	NumResults  int       `xml:"numResults,attr"`
	AdultRating string    `xml:"adultRating,attr"`
	PlaCount    int       `xml:"plaCount,attr"`
	Listings    []Listing `xml:"Listing"`
	NextArgs    string    `xml:"NextArgs"`
}

type Listing struct {
	Rank         int        `xml:"rank,attr"`
	K            string     `xml:"k,attr"`
	AppNs        string     `xml:"appNs,attr"`
	Title        string     `xml:"title,attr"`
	Metadata     string     `xml:"metadata,attr"`
	Description  string     `xml:"description,attr"`
	SiteHost     string     `xml:"siteHost,attr"`
	ImpressionId string     `xml:"ImpressionId,attr"`
	NewEcpi      string     `xml:"new_ecpi,attr"`
	AdultRating  string     `xml:"adultRating,attr"`
	PhoneNumber  string     `xml:"phoneNumber,attr"`
	ClickUrl     ClickUrl   `xml:"ClickUrl"`
	Extensions   Extensions `xml:"Extensions"`
}

type ClickUrl struct {
	Type string `xml:"type,attr"`
	URL  string `xml:",chardata"`
}

type Extensions struct {
	PartnerOptOut   *PartnerOptOut   `xml:"PartnerOptOut"`
	ActionExtension *ActionExtension `xml:"actionExtension"`
}

type PartnerOptOut struct {
	IsAllowed string `xml:"IsAllowed"`
}

type ActionExtension struct {
	ActionItem ActionItem `xml:"actionItem"`
}

type ActionItem struct {
	Text string `xml:"text"`
	Link string `xml:"link"`
}
