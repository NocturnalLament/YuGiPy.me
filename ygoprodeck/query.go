package ygoprodeck

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func URLAttrBuilder(y *YuGiOhProDeckSearchData) string {
	yAttrMap := y.Mapify()
	URLAttrs := url.Values{}
	for k, v := range yAttrMap {
		if v != "Default" && v != "0" {
			URLAttrs.Add(k, v)
		}
	}
	encodedURL := URLAttrs.Encode()
	url := fmt.Sprintf("https://db.ygoprodeck.com/api/v7/cardinfo.php?%s", encodedURL)
	return url
}

func Query(url string) (*CardData, error) {
	// Query the YgoProDeck API
	// Return the response in a YgoProDeckCard struct
	URL := url
	data, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer data.Body.Close()
	var y CardData
	err = json.NewDecoder(data.Body).Decode(&y)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if len(y.Data) > 0 {
		return &y, nil
	}
	noCardErr := fmt.Errorf("no card found")
	return nil, noCardErr
}
