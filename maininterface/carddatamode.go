package maininterface

import (
	"fmt"
	"github.com/Iilun/survey/v2"
	"github.com/NocturnalLament/yugigo/ygoprices"
	"github.com/NocturnalLament/yugigo/ygoprodeck"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

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
