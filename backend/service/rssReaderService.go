package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"rssReader/models"
)

var RssReaderInstance RssReader = RssReader{}

type RssReader struct {}

func (reader *RssReader) Read(provider *models.Provider) (models.Rss2) {
	// http.Get() returns a *http.Response
	response, err := http.Get(provider.URI)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// read the response body
	// io.ReadAll() returns a []byte
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// parse the xml
	// xml.Unmarshal() takes []byte and a pointer to a struct
	var rss models.Rss2
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return rss
}
