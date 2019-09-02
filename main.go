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
	XMLName xml.Name  `xml:"rss"`
	Text    string    `xml:",chardata"`
	Version string    `xml:"version,attr"`
	RSS     []Channel `xml:"channel"`
}

// Channel struct which contains service info
type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Channel []Item   `xml:"item"`
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
	fmt.Printf("%s", byteValue)
	var item RSS

	xml.Unmarshal(byteValue, &item)
	fmt.Println(item.RSS[0].Channel[0])
	for i := 0; i < len(item.Channel); i++ {
		fmt.Println("Title: " + item.Channel[i].Title)
		fmt.Println("Link: " + item.Channel[i].Link)
		fmt.Println("PubDate: " + item.Channel[i].PubDate)
		fmt.Println("Guid: " + item.Channel[i].GUID)
		fmt.Println("Description: " + item.Channel[i].Description)
	}
}
