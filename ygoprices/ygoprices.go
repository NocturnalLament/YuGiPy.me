package ygoprices

import (
	"encoding/json"
	"fmt"
	"github.com/NocturnalLament/yugigo/display"
	"io"
	"log"
	"net/http"
	"net/url"
)

var _ display.CardDataDisplay

type CardCollection struct {
	SearchTerm string
	Cards      []Card `json:"Data"`
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
	fmt.Printf("issue: %s\n", s)
	d, e := http.Get(s)
	if e != nil {
		log.Fatal(e)
	}
	defer d.Body.Close()
	b, e := io.ReadAll(d.Body)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(string(b))
	var y CardCollection
	e = json.Unmarshal(b, &y)
	fmt.Println(e)
	if e != nil {
		return nil, e
	}
	return &y, nil
}

func (c Card) DisplayData() string {
	output := ""
	output += fmt.Sprintf("Name: %s\n", c.Name)
	output += fmt.Sprintf("Print Tag: %s\n", c.PrintTag)
	output += fmt.Sprintf("Rarity: %s\n", c.Rarity)
	output += "Price SearchData:\n"
	output += fmt.Sprintf("Status: %s\n", c.PriceData.Status)
	output += "Prices: \n"
	output += fmt.Sprintf("High: %.2f\n", c.PriceData.Data.Prices.High)
	output += fmt.Sprintf("Low: %.2f\n", c.PriceData.Data.Prices.Low)
	output += fmt.Sprintf("Average: %.2f\n", c.PriceData.Data.Prices.Average)
	output += fmt.Sprintf("Shift: %.2f\n", c.PriceData.Data.Prices.Shift)
	output += fmt.Sprintf("Shift 3: %.2f\n", c.PriceData.Data.Prices.Shift3)
	output += fmt.Sprintf("Shift 7: %.2f\n", c.PriceData.Data.Prices.Shift7)
	output += fmt.Sprintf("Shift 21: %.2f\n", c.PriceData.Data.Prices.Shift21)
	output += fmt.Sprintf("Shift 30: %.2f\n", c.PriceData.Data.Prices.Shift30)
	output += fmt.Sprintf("Shift 90: %.2f\n", c.PriceData.Data.Prices.Shift90)
	output += fmt.Sprintf("Shift 180: %.2f\n", c.PriceData.Data.Prices.Shift180)
	output += fmt.Sprintf("Shift 365: %.2f\n", c.PriceData.Data.Prices.Shift365)
	output += fmt.Sprintf("Updated At: %s\n", c.PriceData.Data.Prices.UpdatedAt)
	return output
}
