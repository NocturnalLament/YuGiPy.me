package main

import (
	"fmt"

	"github.com/Iilun/survey/v2"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/NocturnalLament/yugigo/display"
	"github.com/NocturnalLament/yugigo/ygoprices"
	"github.com/NocturnalLament/yugigo/ygoprodeck"
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

type CardDataDisplay interface {
	DisplayData()
}

type ExecutionMode interface {
	Execute()
}

type CardDataMode struct {
	Data *ygoprodeck.YuGiOhProDeckSearchData
	App  *tview.Application
}
type CardPricesMode struct {
	CardName string
	CardData *ygoprices.Card
	App      *tview.Application
	Flex     *tview.Flex
}
type ServerMode struct{}

func (c *CardDataMode) Execute() {
	data, err := ygoprodeck.GetDataToSearch()
	if err != nil {
		return
	}
	c.Data = data
}

func GetCardDataPrompt() (*ygoprodeck.CardData, error) {
	card, err := ygoprodeck.GetDataToSearch()
	if err != nil {
		return nil, err
	}
	url := ygoprodeck.URLAttrBuilder(card)
	cardData, err := ygoprodeck.Query(url)
	if err != nil {
		return nil, err
	}

	return cardData, nil
}

func SelectCardQuery() (string, int, error) {
	cardData, err := GetCardDataPrompt()
	if err != nil {
		return "", -1, err
	}
	names := cardData.GetCardNames()
	nameToSearchSelect := survey.Select{
		Message: "Select a card to search for:",
		Options: names,
	}
	var nameToSearch string
	survey.AskOne(&nameToSearchSelect, &nameToSearch)
	fmt.Println(nameToSearch)
	prices, err := ygoprices.QueryPrices(nameToSearch)
	if err != nil {
		fmt.Printf("error querying prices: %v", err)
	}
	amountOfCards := len(prices.Cards)
	fmt.Println(amountOfCards)
	if amountOfCards == 0 {
		return "", -1, fmt.Errorf("no cards found")
	}
	return nameToSearch, amountOfCards, nil
}

func (c *CardPricesMode) setupInputCapture(amountOfCards int, prices *ygoprices.CardCollection) {
	cardIndex := 0
	c.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			c.App.Stop()

		case tcell.KeyEnter:
			if cardIndex+1 <= amountOfCards {
				cardIndex++
				c.Flex.Clear()
				display.DisplayCardQueryData(c.App, c.Flex, prices.Cards[cardIndex])
			} else {
				return event
			}
		}
		return event
	})
}

func (c *CardPricesMode) SetupView(prices *ygoprices.CardCollection, amountOfCards int) {
	c.App = tview.NewApplication()
	cardIndex := 0
	c.Flex = tview.NewFlex().SetDirection(tview.FlexRow)
	c.setupInputCapture(len(prices.Cards), prices)
	display.DisplayCardQueryData(c.App, c.Flex, prices.Cards[cardIndex])

}

func (c *CardPricesMode) Execute() {
	nameToSearch, amountOfCards, err := SelectCardQuery()
	if amountOfCards == -1 {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
	}
	prices, err := ygoprices.QueryPrices(nameToSearch)
	if err != nil {
		fmt.Println(err)
	}
	//Begin View Logic.

	c.SetupView(prices, amountOfCards)
	c.App.Run()

}

func ModeSwitch(mode string) ExecutionMode {
	m := ExecutionMode(nil)
	switch mode {
	case "Card Data":
		m = &CardDataMode{}
	case "Card Prices":
		m = &CardPricesMode{}
	}
	return m
}

func PickMode() string {
	modes := []string{"Card Data", "Card Prices", "Server"}
	prompt := survey.Select{
		Message: "Select a mode to run in:",
		Options: modes,
	}
	var mode string
	survey.AskOne(&prompt, &mode)
	return mode
}

func main() {
	//Get data to search
	mode := PickMode()
	m := ModeSwitch(mode)
	m.Execute()
	//Build URL
	//url := ygoprodeck.URLAttrBuilder(data)
	//Query the API
	//cardData, err := ygoprodeck.Query(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//Print the card data
	//Get result to display
	//names := cardData.GetCardNames()
	//resToDisplay := ygoprodeck.GetResultToDisplay(names)
	//fmt.Println(resToDisplay)
	/* prices, err := ygoprices.QueryPrices(resToDisplay)
	if err != nil {
		fmt.Println(err)
	} */
	/*val := fmt.Sprintf("Returns: %v", len(prices.Cards))
	fmt.Println(val) */
	/*app := tview.NewApplication()
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	cardIndex := 0
	//cardData.DisplayCard(app, flex)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			flex.Clear()
			cardData.DisplayCard(app, flex, cardIndex)
		case tcell.KeyTAB:
			cardIndex++
			flex.Clear()
			cardData.DisplayCard(app, flex, cardIndex)
		case tcell.KeyEscape:
			app.Stop()
		}
		return event
	})
	cardData.DisplayCard(app, flex, cardIndex) */
}
