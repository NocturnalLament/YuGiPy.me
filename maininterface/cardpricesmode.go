package maininterface

import (
	"database/sql"
	"fmt"
	"github.com/Iilun/survey/v2"
	"github.com/NocturnalLament/yugigo/display"
	"github.com/NocturnalLament/yugigo/ygoprices"
	"github.com/NocturnalLament/yugigo/ygoprodeck"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CardPricesMode struct {
	cMode        SubmodeOperator
	CardName     string
	SetName      string
	CardData     *ygoprices.Card
	App          *tview.Application
	Flex         *tview.Flex
	Data         *ygoprices.YgoPricesCardData
	cardSelected bool
	CardUrl      string
	NewPrice     bool
	Prices       []CardPricesMode
}

func (c *CardPricesMode) setCMode(mode SubmodeOperator) {
	c.cMode = mode
	c.modeSwitch()
}

func (c CardPricesMode) modeSwitch() {
	switch c.cMode {
	case Read:
		success, err := c.ReadData()
		if err != nil {
			fmt.Println(err)
			return
		}
		if success {
			fmt.Println("Data read successfully")
		}
	}
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
		CardUrl:      "",
		NewPrice:     false,
		Prices:       nil,
	}
}

func (c *CardPricesMode) Read() error {
	db, err := sql.Open("sqlite3", "card_data.db")
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT * FROM card_data")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var cData CardTrackingData
		err = cData.LoadSql(rows)
	}
	return nil
}

func (c *CardPricesMode) LoadSql(rows *sql.Rows) error {

	return nil
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
			CPricesMode.CardUrl = ygoprices.QueryURLBuilder(CPricesMode.CardName)
			y := ygoprices.YgoPricesCardData{
				CardName:        card.Name,
				PrintTag:        card.PrintTag,
				CardPrice:       ygoprices.YGOCardPrice(card.PriceData.Data.Prices.Average),
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
			CPricesMode.Data = &y
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

func (c *CardPricesMode) initializeMode() {
	options := []string{"Read", "Write", "Load"}
	prompt := survey.Select{
		Message: "Read/Write/Load Data?",
		Options: options,
	}
	var response string
	if err := survey.AskOne(&prompt, &response); err != nil {
		fmt.Println(err)
	}
	if response == "Write" {
		c.NewPrice = true
	} else if response == "Read" {
		c.NewPrice = false
		c.setCMode(Read)
	} else if response == "Load" {
		return
	}
}

func (c *CardPricesMode) ReadData() (bool, error) {
	//db, err := sql.Open("sqlite3", "card_data.db")
	//fmt.Println("Got here")
	//if err != nil {
	//	return false, err
	//}
	//defer db.Close()
	//rows, err := db.Query("SELECT CardName, CardSetName, PrintTag, CardPrice, High, Low, Average, Shift, Shift3, Shift7, Shift21, Shift30, Shift90, Shift180, Shift365, TimeLastUpdated FROM card_data")
	////Average, Shift, Shift3, Shift7, Shift21, Shift90, Shift180, Shift365, TimeLastUpdated, ImageUrl, CardURL, TrackedTime
	//if err != nil {
	//	return false, err
	//}
	//defer rows.Close()
	//
	////for rows.Next() {
	////	card := NewCPricesMode()
	////	err = rows.Scan(&c.CardName, &c.SetName, &c.CardData.PrintTag, &c.Data.CardPrice, &c.Data.High, &c.Data.Low,
	////		&c.Data.Average, &c.Data.Shift, &c.Data.Shift3, &c.Data.Shift7, &c.Data.Shift21, &c.Data.Shift30, &c.Data.Shift90,
	////		&c.Data.Shift180, &c.Data.Shift365, &c.Data.TimeLastUpdated, &c.Data.ImageUrl, &c.CardUrl, &c.Data.TrackedTime)
	////	if err != nil {
	////		fmt.Println(err)
	////		return false, err
	////	}
	////cCard := NewCPricesMode()
	//for rows.Next() {
	//	var CardName string
	//	var CardSetPrice string
	//	var PrintTag string
	//	var CardPrice float64
	//	var High float64
	//	var Low float64
	//	var Average float64
	//	var Shift float64
	//	var Shift3 float64
	//	var Shift7 float64
	//	var Shift21 float64
	//	var Shift30 float64
	//	var Shift90 float64
	//	var Shift180 float64
	//	var Shift365 float64
	//	var TimeLastUpdated string
	//	e := rows.Scan(&CardName, &CardSetPrice, &PrintTag, &CardPrice, &High, &Low, &Average, &Shift, &Shift3, &Shift7, &Shift21, &Shift30, &Shift90, &Shift180, &Shift365, &TimeLastUpdated)
	//	if e != nil {
	//		return false, e
	//	}
	//	fmt.Println(CardName)
	//	fmt.Println(CardSetPrice)
	//	fmt.Println(PrintTag)
	//}
	cardStruct := ygoprices.NewYgoPriceData()
	stuff, err := cardStruct.ReadData()
	if err != nil {
		return false, nil
	}
	for _, card := range stuff {
		fmt.Println(card.CardName)
	}
	return false, nil
}

func (c *CardPricesMode) WriteData() (bool, error) {
	return false, nil
}

func (c *CardPricesMode) Execute() {
	c.initializeMode()
	if c.NewPrice {
		CPricesMode = NewCPricesMode()
		nameToSearch, prices, amountOfCards, err := SelectCardQuery()
		CPricesMode.CardName = nameToSearch
		fmt.Println(nameToSearch)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(nameToSearch)

		pricesData := formatDataForOutput(prices)
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
	} else {
		fmt.Println("Exiting...")
	}
}

func (c CardPricesMode) Insert() (bool, error) {
	db, err := sql.Open("sqlite3", "card_data.db")
	if err != nil {
		return false, err
	}
	sqliteInsertStatement := `INSERT INTO card_data (CardName, CardSetName, PrintTag, CardPrice, High, Low, Average, Shift, Shift3, Shift7, Shift21, Shift30, Shift90, Shift180, Shift365, TimeLastUpdated, ImageUrl, CardURL, TrackedTime) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	statement, err := db.Prepare(sqliteInsertStatement)
	if err != nil {
		return false, err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(statement)
	_, err = statement.Exec(c.CardName, c.Data.CardName, c.Data.PrintTag, c.Data.CardPrice, c.Data.High, c.Data.Low,
		c.Data.Average, c.Data.Shift, c.Data.Shift3, c.Data.Shift7, c.Data.Shift21, c.Data.Shift30, c.Data.Shift90,
		c.Data.Shift180, c.Data.Shift365, c.Data.TimeLastUpdated, c.Data.ImageUrl, c.CardUrl, c.Data.TrackedTime)
	if err != nil {
		return false, err
	}
	return true, nil
}
