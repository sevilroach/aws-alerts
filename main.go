package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//RSS struct which contains the whole status data
type RSS struct {
	XMLName  xml.Name  `xml:"rss"`
	Text     string    `xml:",chardata"`
	Version  string    `xml:"version,attr"`
	Channels []Channel `xml:"channel"`
}

// Channel struct which contains service info
type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Item   `xml:"item"`
}

// Item struct which contains status info of the service
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	GUID        string   `xml:"guid"`
	Description string   `xml:"description"`
}

func main() {
	// Amazon Elastic Compute Cloud (N. Virginia) Service Status
	res, err := http.Get("https://status.aws.amazon.com/rss/ec2-us-east-1.rss")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var rssData RSS

	xml.Unmarshal(byteValue, &rssData)

	for _, item := range rssData.Channels[0].Items {
		fmt.Println("Title: " + item.Title)
		fmt.Println("Link: " + item.Link)
		fmt.Println("PubDate: " + item.PubDate)
		fmt.Println("Guid: " + item.GUID)
		fmt.Println("Description: " + item.Description)
		fmt.Println("---------------------------------------")
	}
}
