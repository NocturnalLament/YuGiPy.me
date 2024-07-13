package ygoprices

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CardCollection struct {
	Cards []Card `json:"Data"`
}

type Card struct {
	Name      string    `json:"name"`
	PrintTag  string    `json:"print_tag"`
	Rarity    string    `json:"rarity"`
	PriceData PriceData `json:"price_data"`
}

type PriceData struct {
	Status string `json:"status"`
	Data   struct {
		Listings []interface{} `json:"listings"` // Assuming listings is an array of unknown objects or empty
		Prices   PriceDetails  `json:"prices"`
	} `json:"data"`
}

type PriceDetails struct {
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Average   float64 `json:"average"`
	Shift     float64 `json:"shift"`
	Shift3    float64 `json:"shift_3"`
	Shift7    float64 `json:"shift_7"`
	Shift21   float64 `json:"shift_21"`
	Shift30   float64 `json:"shift_30"`
	Shift90   float64 `json:"shift_90"`
	Shift180  float64 `json:"shift_180"`
	Shift365  float64 `json:"shift_365"`
	UpdatedAt string  `json:"updated_at"`
}

func QueryPrices(cardName string) (*CardCollection, error) {
	// Query the YugiohPrices API
	// Return the response in a CardCollection struct
	encodedCardName := url.QueryEscape(cardName)
	s := fmt.Sprintf("http://yugiohprices.com/api/get_card_prices/%s", encodedCardName)
	d, e := http.Get(s)
	if e != nil {
		log.Fatal(e)
	}
	defer d.Body.Close()
	b, e := io.ReadAll(d.Body)
	if e != nil {
		log.Fatal(e)
	}
	//fmt.Println(string(b))
	var y CardCollection
	e = json.Unmarshal(b, &y)
	if e != nil {
		return nil, e
	}
	return &y, nil
}
