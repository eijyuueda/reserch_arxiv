package pkg

import (
	"encoding/xml"
)

type Entry struct {
	ID        string `xml:"id"`
	Title     string `xml:"title"`
	Summary   string `xml:"summary"`
	Published string `xml:"published"`
	Links     []struct {
		Rel  string `xml:"rel,attr"`
		HRef string `xml:"href,attr"`
	} `xml:"link"`
}

type Feed struct {
	Entries []Entry `xml:"entry"`
}

// XML形式にパースする
func ParseXML(body []byte) (Feed, error) {
	var feed Feed
	err := xml.Unmarshal(body, &feed)
	return feed, err
}
