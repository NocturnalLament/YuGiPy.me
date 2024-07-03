package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Iilun/survey/v2"
	"github.com/manifoldco/promptui"
)

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

type YugiohProDeckSearchByType int

type YuGiOhProDeckStructFields map[string]string

func InitialzeYuGiOhProDeckMap() YuGiOhProDeckStructFields {
	return YuGiOhProDeckStructFields{
		"Name":        "name",
		"FName":       "fname",
		"Id":          "id",
		"KonamiId":    "id",
		"Type":        "type",
		"Atk":         "atk",
		"Def":         "def",
		"Level":       "level",
		"Race":        "race",
		"Attribute":   "attribute",
		"Link":        "link",
		"LinkMarkers": "linkmarkers",
		"Scale":       "scale",
		"CardSet":     "cardset",
		"Archetype":   "archetype",
		"Banlist":     "banlist",
		"Sort":        "sort",
		"Format":      "format",
		"Misc":        "misc",
	}
}

type YuGiOhProDeckSearchData struct {
	Name        string
	FName       string
	Id          int
	KonamiId    int
	Type        string
	Atk         int
	Def         int
	Level       int
	Race        string
	Attribute   string
	Link        string
	LinkMarkers []string
	Scale       int
	CardSet     string
	Archetype   string
	Banlist     string
	Sort        string
	Format      string
	Misc        bool // Will either be unpassed or if true will be passed as "yes"
}

func GetBanList() string {
	banLists := []string{"TCG", "OCG", "GOAT"}
	prompt := survey.Select{
		Message: "Select the banlist",
		Options: banLists,
	}
	selected := ""
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	return selected
}

func GetCardAttributes() string {
	cardAttributes := []string{"Dark", "Divine", "Earth", "Fire", "Light", "Water", "Wind"}
	prompt := survey.Select{
		Message: "Select the card attribute",
		Options: cardAttributes,
	}
	selected := ""
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	return selected
}

func GetLinkMarkers() string {
	linkMarkersOptions := []string{"Top", "Bottom", "Left", "Right", "Bottom-Left", "Bottom-Right", "Top-Left", "Top-Right"}
	prompt := survey.MultiSelect{
		Message:  "Select your link markers",
		Options:  linkMarkersOptions,
		PageSize: 8,
	}
	selected := []string{}
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.Join(selected, ",")
}

func ProDeckPrompt() {
	structFields := []string{"Name", "Fuzzy Name", "ID", "Konami ID", "Type", "ATK", "DEF", "Level", "Race", "Attribute", "Link", "LinkMarkers", "Scale", "Card Set", "Archetype", "Banlist", "Sort", "Format", "Misc"}
	prompt := survey.MultiSelect{
		Message:  "Select the fields you want to include in your search query",
		Options:  structFields,
		PageSize: 10,
	}
	selected := []string{}
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(selected)
	//return nil, nil
}

func GetDataToSearch() (*YuGiOhProDeckSearchData, error) {

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return fmt.Errorf("Invalid input")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Enter the card ID",
		Validate: validate,
	}
	res, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	return nil, nil
}

func main() {
	// Call the function
	fmt.Println("Hello, World!")
	ProDeckPrompt()
	fmt.Println(GetLinkMarkers())
	fmt.Println(GetCardAttributes())
	fmt.Println(GetBanList())

}
