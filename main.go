package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/NocturnalLament/yugigo/ygoprodeck"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

// http://yugiohprices.com/api/get_card_prices/card_name/print_tag
type YugiohPricesDataByCardPrintTag struct {
	Status string `json:"status"`
	Data   []struct {
		Name      string `json:"name"`
		PrintTag  string `json:"print_tag"`
		Rarity    string `json:"rarity"`
		PriceData []struct {
			Status string `json:"status"`
			Data   []struct {
				Prices []struct {
					High      string `json:"high"`
					Low       string `json:"low"`
					Average   string `json:"average"`
					Shift     int    `json:"shift"`
					Shift3    int    `json:"shift3"`
					Shift7    int    `json:"shift7"`
					Shift30   int    `json:"shift30"`
					Shift90   int    `json:"shift90"`
					Shift180  int    `json:"shift180"`
					Shift365  int    `json:"shift365"`
					UpdatedAt string `json:"updated_at"`
				} `json:"prices"`
			} `json:"data"`
		} `json:"price_data"`
	} `json:"data"`
}

// http://yugiohprices.com/api/get_card_prices/card_name/print_tag/rarity
type YugiohPriceHistorySpecificTagAndRarity struct {
	Status string `json:"status"`
	Data   []struct {
		PriceAverage float32 `json:"price_average"`
		PriceShift   float64 `json:"price_shift"`
		CreatedAt    string  `json:"created_at"`
	} `json:"data"`
}

// http://yugiohprices.com/api/get_card_prices/set_data/{set_name}
type YugioPriceSetData struct {
	Status string `json:"status"`
	Data   []struct {
		Rarities struct {
			Rare         int `json:"Rare"`
			Common       int `json:"Common"`
			SuperRare    int `json:"Super Rare"`
			SecretRare   int `json:"Secret Rare"`
			UltraRare    int `json:"Ultra Rare"`
			UltimateRare int `json:"Ultimate Rare"`
		}
		Average            float32 `json:"average"`
		Lowest             float32 `json:"lowest"`
		Highest            float32 `json:"highest"`
		tcg_booster_values struct {
			High    float32 `json:"high"`
			Low     float32 `json:"low"`
			Average float32 `json:"average"`
		}
		Cards []struct {
			Name    string `json:"name"`
			Numbers []struct {
				Name      string `json:"name"`
				PrintTag  string `json:"print_tag"`
				Rarity    string `json:"rarity"`
				PriceData struct {
					Status string `json:"status"`
					Data   struct {
						Prices struct {
							High      float32 `json:"high"`
							Low       float32 `json:"low"`
							Average   float32 `json:"average"`
							Shift     int     `json:"shift"`
							Shift3    int     `json:"shift3"`
							Shift7    int     `json:"shift7"`
							Shift21   int     `json:"shift21"`
							Shift30   int     `json:"shift30"`
							Shift90   int     `json:"shift90"`
							Shift180  int     `json:"shift180"`
							Shift365  int     `json:"shift365"`
							UpdatedAt string  `json:"updated_at"`
						} `json:"prices"`
					} `json:"data"`
				} `json:"price_data"`
			} `json:"numbers"`
			CardType    string `json:"card_type"`
			Family      string `json:"family"`
			MonsterType string `json:"type"`
		} `json:"cards"`
	} `json:"data"`
}

//func LogToFile()

func GetDataToSearch() (*ygoprodeck.YuGiOhProDeckSearchData, error) {

	items := ygoprodeck.ProDeckPrompt()
	vals := ygoprodeck.GetValsFromPrompt(items)
	item := vals.ProcessPrompts()
	if item == nil {
		return nil, fmt.Errorf("error processing prompts")
	}
	spew.Dump(item)
	fmt.Printf("Item: %v\n", item.Name)
	return item, nil
}

/* func main() {
	hello := logger.StandardLogger{
		Logger:      log.New(os.Stdout, "Hello: ", log.Ldate|log.Ltime|log.Lshortfile),
		Level:       logger.Info,
		LogFilePath: "logr.log",
	}
	hello.Info("Hello, World!")
	// Call the function
	item, _ := GetDataToSearch()
	mapItem := item.Mapify()
	values := url.Values{}
	for k, v := range mapItem {
		if v != "Default" && v != "0" {
			values.Add(k, v)
		}
	}
	encodedStrin := values.Encode()
	urlThing := "https://db.ygoprodeck.com/api/v7/cardinfo.php?" + encodedStrin
	fmt.Println(urlThing)
	data, err := http.Get(urlThing)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()
	var result map[string]interface{}

	// Decode the response body to the struct or map
	err = json.NewDecoder(data.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	//var y ygoprodeck.YgoProDecData
	//bodyByte, err := io.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	// err = json.Unmarshal(bodyByte, &y)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(y.Data[0].Name)

	// Now `result` holds the decoded JSON data
	var cards []ygoprodeck.YgoProDeckCard

	bodyByte, err := io.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyByte, &cards)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.Body)
	fmt.Println(bodyByte)
	data.Body.Close()
	log := logrus.New()
	file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	log.SetLevel(logrus.DebugLevel)

} */

func main() {
	log := logrus.New()
	file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Out = file
	log.SetLevel(logrus.DebugLevel)

	item, err := GetDataToSearch()
	if err != nil {
		log.Fatal(err)
	}
	mapItem := item.Mapify()
	values := url.Values{}
	for k, v := range mapItem {
		if v != "Default" && v != "0" {
			values.Add(k, v)
		}
	}
	encodedString := values.Encode()
	urlThing := "https://db.ygoprodeck.com/api/v7/cardinfo.php?" + encodedString
	fmt.Println(urlThing)
	data, err := http.Get(urlThing)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	bodyByte, err := io.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	var dataCard ygoprodeck.CardData
	jsonErr := json.Unmarshal(bodyByte, &dataCard)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	spew.Dump(dataCard)
	fmt.Println(dataCard.Data[0].Name)
	//fmt.Println(string(bodyByte))
	// var cards []ygoprodeck.YgoProDeckCard
	// err = json.Unmarshal(bodyByte, &cards)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//spew.Dump(cards)
	//fmt.Println(string(bodyByte))
	//fmt.Println(cards[0].Name)
	// Use `cards` for further processing
}
