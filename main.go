package main

import "fmt"

// https://db.ygoprodeck.com/api/v7/cardinfo.php
type YgoProDecData struct {
	Data struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		CardType    string `json:"type"`
		FrameType   string `json:"frameType"`
		Description string `json:"desc"`
		Atk         int    `json:"atk"`
		Def         int    `json:"def"`
		Level       int    `json:"level"`
		Race        string `json:"race"`
		Attribute   string `json:"attribute"`
		CardSets    []struct {
			SetName   string `json:"set_name"`
			SetCode   string `json:"set_code"`
			SetRarity string `json:"set_rarity"`
			SetPrice  string `json:"set_price"`
		} `json:"card_sets"`
		CardImages []struct {
			Id              int    `json:"id"`
			ImageURL        string `json:"image_url"`
			ImageURLSmall   string `json:"image_url_small"`
			ImageURLCropped string `json:"image_url_cropped"`
		} `json:"card_images"`
		CardPrices []struct {
			CardMarketPrice   string `json:"cardmarket_price"`
			TCGPlayerPrice    string `json:"tcgplayer_price"`
			EbayPrice         string `json:"ebay_price"`
			AmazonPrice       string `json:"amazon_price"`
			CoolstuffincPrice string `json:"coolstuffinc_price"`
		} `json:"card_prices"`
	} `json:"data"`
}

// http://yugiohprices.com/api/get_card_prices/card_name
type YugiohPricesByCardName struct {
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

type YugiohPriceHistorySpecificTagAndRarity struct {
	Status string `json:"status"`
	Data   []struct {
		PriceAverage float32 `json:"price_average"`
		PriceShift   float64 `json:"price_shift"`
		CreatedAt    string  `json:"created_at"`
	} `json:"data"`
}

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
					data   struct {
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

func main() {
	// Call the function
	fmt.Println("Hello, World!")
}
