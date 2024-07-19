package main

import (
	"fmt"
	"github.com/Iilun/survey/v2"
	"github.com/NocturnalLament/yugigo/display"
	"github.com/NocturnalLament/yugigo/ygoprices"
	"github.com/NocturnalLament/yugigo/ygoprodeck"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// YugiohPricesDataByCardPrintTag http://yugiohprices.com/api/get_card_prices/card_name/print_tag
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

// YugiohPriceHistorySpecificTagAndRarity http://yugiohprices.com/api/get_card_prices/card_name/print_tag/rarity
type YugiohPriceHistorySpecificTagAndRarity struct {
	Status string `json:"status"`
	Data   []struct {
		PriceAverage float32 `json:"price_average"`
		PriceShift   float64 `json:"price_shift"`
		CreatedAt    string  `json:"created_at"`
	} `json:"data"`
}

// YugioPriceSetData http://yugiohprices.com/api/get_card_prices/set_data/{set_name}
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
		Average          float32 `json:"average"`
		Lowest           float32 `json:"lowest"`
		Highest          float32 `json:"highest"`
		tcgBoosterValues struct {
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

type CarddataModeOutputStorage struct {
	CardName string
	CardData *CardDataMode
}

func NewCardDataModeOutputStorage() *CarddataModeOutputStorage {
	return &CarddataModeOutputStorage{
		CardName: "",
		CardData: nil,
	}
}

type CardDataMode struct {
	SearchData       *ygoprodeck.YuGiOhProDeckSearchData
	ReturnedCardData *ygoprodeck.CardData
	App              *tview.Application
	Flex             *tview.Flex
	CardSelected     bool
}

type CardPricesMode struct {
	CardName     string
	SetName      string
	CardData     *ygoprices.Card
	App          *tview.Application
	Flex         *tview.Flex
	Data         *ygoprices.YgoPricesCardData
	cardSelected bool
}

var CDataMode *CardDataMode

var CPricesMode *CardPricesMode

func NewCPricesMode() *CardPricesMode {
	return &CardPricesMode{
		CardName:     "",
		SetName:      "",
		CardData:     nil,
		App:          nil,
		Flex:         nil,
		Data:         nil,
		cardSelected: false,
	}
}

func NewCDataMode() *CardDataMode {
	return &CardDataMode{
		SearchData:       nil,
		ReturnedCardData: nil,
		App:              nil,
		Flex:             nil,
	}
}

type ServerMode struct{}

func (c *CardDataMode) SetupInputCapture(cardIndex int, amountOfCards int, cardData *ygoprodeck.CardData, app *tview.Application, flex *tview.Flex) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyEnter:
			if cardIndex < amountOfCards {
				cardIndex++
				flex.Clear()
				cardData.DisplayCard(app, flex, cardIndex)
			} else {
				flex.Clear()
				app.Stop()
				return event
			}
		default:
			switch event.Rune() {
			case 's':
				flex.Clear()
			}
		}
		return event
	})
	if amountOfCards > 1 {
		cardData.DisplayCard(app, flex, cardIndex)
	}
}

func (c *CardDataMode) Execute() {
	CDataMode = NewCDataMode()
	data, err := ygoprodeck.GetDataToSearch()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.SearchData = data
	CDataMode.SearchData = data
	url := ygoprodeck.URLAttrBuilder(data)
	cardData, err := ygoprodeck.Query(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	CDataMode.ReturnedCardData = cardData
	app := tview.NewApplication()
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	amountOfCards := len(cardData.Data)
	cardIndex := 0
	c.SetupInputCapture(cardIndex, amountOfCards, cardData, app, flex)
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

func SelectCardQuery() (string, *ygoprices.CardCollection, int, error) {
	cardData, err := GetCardDataPrompt()
	if err != nil {
		return "", nil, -1, err
	}
	names := cardData.GetCardNames()
	nameToSearchSelect := survey.Select{
		Message: "Select a card to search for:",
		Options: names,
	}
	var nameToSearch string
	if err = survey.AskOne(&nameToSearchSelect, &nameToSearch); err != nil {
		return "", nil, -1, err
	}
	fmt.Println(nameToSearch)
	prices, err := ygoprices.QueryPrices(nameToSearch)
	fmt.Println(prices)
	if err != nil {
		fmt.Printf("error querying prices: %v", err)
		return "", nil, -1, err
	}
	amountOfCards := len(prices.Cards)
	fmt.Println(amountOfCards)
	if amountOfCards == 0 {
		return "", nil, -1, fmt.Errorf("no cards found")
	}
	return nameToSearch, prices, amountOfCards, nil
}

func (c *CardPricesMode) setupInputCapture(amountOfCards int, prices *ygoprices.CardCollection) {
	cardIndex := 0
	c.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			c.App.Stop()

		case tcell.KeyEnter:
			if c.cardSelected {
				c.Flex.Clear()
				display.DisplayEndOfPrices(c.App, c.Flex)
				c.App.Stop()
			} else if cardIndex < amountOfCards {
				c.Flex.Clear()
				display.DisplayCardQueryData(c.App, c.Flex, len(prices.Cards), cardIndex, prices.Cards[cardIndex])
				cardIndex++

			} else if cardIndex == amountOfCards {
				c.Flex.Clear()
				display.DisplayEndOfPrices(c.App, c.Flex)
				c.App.Stop()
			}
		case tcell.KeyTAB:
			c.Flex.Clear()
			c.cardSelected = true
			fmt.Println("Tab pressed")
			selectedIndex := 0
			if cardIndex > 0 {
				selectedIndex = cardIndex - 1
			}
			card := prices.Cards[selectedIndex]
			CPricesMode.CardData = &card
			CPricesMode.SetName = card.Name
			y := ygoprices.YgoPricesCardData{
				CardName:        card.Name,
				PrintTag:        card.PrintTag,
				CardPrice:       card.PriceData.Data.Prices.Average,
				High:            card.PriceData.Data.Prices.High,
				Low:             card.PriceData.Data.Prices.Low,
				Average:         card.PriceData.Data.Prices.Average,
				Shift:           card.PriceData.Data.Prices.Shift,
				Shift3:          card.PriceData.Data.Prices.Shift3,
				Shift7:          card.PriceData.Data.Prices.Shift7,
				Shift21:         card.PriceData.Data.Prices.Shift21,
				Shift30:         card.PriceData.Data.Prices.Shift30,
				Shift90:         card.PriceData.Data.Prices.Shift90,
				Shift180:        card.PriceData.Data.Prices.Shift180,
				Shift365:        card.PriceData.Data.Prices.Shift365,
				TimeLastUpdated: card.PriceData.Data.Prices.UpdatedAt,
			}
			display.DisplaySelectedCardPrice(c.App, c.Flex, y.CardString())

		default:
			switch event.Rune() {
			case 's':

				c.Flex.Clear()

			}
			return event
		}
		return event
	})
}

func (c *CardPricesMode) SetupView(prices *ygoprices.CardCollection, amountOfCards int) {
	c.App = tview.NewApplication()
	cardIndex := 0
	c.Flex = tview.NewFlex().SetDirection(tview.FlexRow)
	c.setupInputCapture(len(prices.Cards), prices)
	display.DisplayCardQueryData(c.App, c.Flex, len(prices.Cards), cardIndex, prices.Cards[cardIndex])

	if err := c.App.SetRoot(c.Flex, true).SetFocus(c.Flex).Run(); err != nil {
		panic(err)
	}
}

func (c *CardPricesMode) Execute() {
	CPricesMode = NewCPricesMode()
	nameToSearch, prices, amountOfCards, err := SelectCardQuery()
	CPricesMode.CardName = nameToSearch
	fmt.Println(nameToSearch)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nameToSearch)

	pricesData := []*ygoprices.YgoPricesCardData{}
	for _, card := range prices.Cards {
		priceDataStruct := ygoprices.NewYgoPriceData()
		priceDataStruct.CardName = card.Name
		priceDataStruct.PrintTag = card.PrintTag
		priceDataStruct.CardPrice = card.PriceData.Data.Prices.Average
		priceDataStruct.High = card.PriceData.Data.Prices.High
		priceDataStruct.Low = card.PriceData.Data.Prices.Low
		priceDataStruct.Average = card.PriceData.Data.Prices.Average
		priceDataStruct.Shift = float64(card.PriceData.Data.Prices.Shift)
		priceDataStruct.Shift3 = float64(card.PriceData.Data.Prices.Shift3)
		priceDataStruct.Shift7 = float64(card.PriceData.Data.Prices.Shift7)
		priceDataStruct.Shift21 = float64(card.PriceData.Data.Prices.Shift21)
		priceDataStruct.Shift30 = float64(card.PriceData.Data.Prices.Shift30)
		priceDataStruct.Shift90 = float64(card.PriceData.Data.Prices.Shift90)
		priceDataStruct.Shift180 = float64(card.PriceData.Data.Prices.Shift180)
		priceDataStruct.Shift365 = card.PriceData.Data.Prices.Shift365
		pricesData = append(pricesData, priceDataStruct)
	}
	fmt.Printf("Returned: %d\n", len(pricesData))

	//Begin View Logic.

	c.SetupView(prices, amountOfCards)
	if err = c.App.Run(); err != nil {
		fmt.Println(err)
		return
	}
	c.App.Stop()
	c.Flex.Clear()
	fmt.Println("Hello world!")
	fmt.Println(CPricesMode.CardName)
	fmt.Println(CPricesMode.SetName)
	d := ygoprodeck.YuGiOhProDeckSearchData{
		Name:    CPricesMode.CardName,
		CardSet: CPricesMode.SetName,
	}
	url := ygoprodeck.URLAttrBuilder(&d)
	fmt.Println(url)
	newCardData, err := ygoprodeck.Query(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newCardData.DisplayData())
}

func ModeSwitch(mode string) ExecutionMode {
	m := ExecutionMode(nil)
	switch mode {
	case "Card SearchData":
		m = &CardDataMode{}
	case "Card Prices":
		m = &CardPricesMode{}

	}
	return m
}

func PickMode() string {
	modes := []string{"Card SearchData", "Card Prices", "Server"}
	prompt := survey.Select{
		Message: "Select a mode to run in:",
		Options: modes,
	}
	var mode string
	if err := survey.AskOne(&prompt, &mode); err != nil {
		fmt.Println(err)
	}
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
