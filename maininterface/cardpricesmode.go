package maininterface

import (
	"database/sql"
	"fmt"
	survey "github.com/Iilun/survey/v2"
	"github.com/NocturnalLament/yugigo/display"
	"github.com/NocturnalLament/yugigo/displaymanager"
	"github.com/NocturnalLament/yugigo/ygoprices"
	"github.com/NocturnalLament/yugigo/ygoprodeck"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CardPricesMode struct {
	ProgramSubmode
	cMode                SubmodeOperator
	CardName             string
	SetName              string
	CardIndex            int
	CardData             *ygoprices.Card
	Display              *displaymanager.DisplayManager
	DisplaySetupCallback func()
	Data                 *ygoprices.YgoPricesCardData
	cardSelected         bool
	CardUrl              string
	NewPrice             bool
	CardSelected         bool
	Prices               []CardPricesMode
	Collection           *ygoprices.CardCollection
	LoadData             *ygoprices.YgoPricesCardData
}

func (c *CardPricesMode) setCMode(mode SubmodeOperator) {
	c.cMode = mode
	c.modeSwitch()
}

func (c *CardPricesMode) modeSwitch() {
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

func NewCPricesMode() *CardPricesMode {
	return &CardPricesMode{
		CardName: "",
		SetName:  "",
		CardData: nil,
		Display: &displaymanager.DisplayManager{
			App:  tview.NewApplication(),
			Flex: tview.NewFlex().SetDirection(tview.FlexRow),
		},
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
	//defer rows.Close()
	func(rows *sql.Rows) {
		errClose := rows.Close()
		if errClose != nil {
			fmt.Println(errClose)
		}
	}(rows)
	for rows.Next() {
		var cData CardTrackingData
		err = cData.LoadSql(rows)
	}
	return nil
}

func (c *CardPricesMode) SetupInputCapture(amountOfCards int, prices *ygoprices.CardCollection) {
	if c.Display == nil {
		fmt.Println("Display nil")
		c.Display = &displaymanager.DisplayManager{
			App:  tview.NewApplication(),
			Flex: tview.NewFlex().SetDirection(tview.FlexRow),
		}

	}

	c.Display.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			c.Display.App.Stop()

		case tcell.KeyEnter:
			if c.CardSelected == true {
				_, err := c.Insert()
				if err != nil {
					fmt.Println(err)
				}
			}
			if c.cardSelected {
				c.Display.Flex.Clear()
				display.DisplayEndOfPrices(c.Display.App, c.Display.Flex)
				c.Display.App.Stop()
			} else if c.CardIndex < amountOfCards {
				c.Display.Flex.Clear()
				display.DisplayCardQueryData(c.Display.App, c.Display.Flex, len(prices.Cards), c.CardIndex, prices.Cards[c.CardIndex])
				c.CardIndex++

			} else if c.CardIndex == amountOfCards {
				c.Display.Flex.Clear()
				display.DisplayEndOfPrices(c.Display.App, c.Display.Flex)
				c.Display.App.Stop()
			}
		case tcell.KeyTAB:
			c.Display.Flex.Clear()
			c.cardSelected = true
			fmt.Println("Tab pressed")
			selectedIndex := 0
			if c.CardIndex > 0 {
				selectedIndex = c.CardIndex - 1
			}
			card := prices.Cards[selectedIndex]
			c.CardData = &card
			c.SetName = card.Name
			c.CardUrl = ygoprices.QueryURLBuilder(c.CardName)

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
			c.Data = &y

			c.CardSelected = true
			display.DisplaySelectedCardPrice(c.Display.App, c.Display.Flex, y.CardString())

		default:
			switch event.Rune() {
			case 's':

				c.Display.Flex.Clear()

			}
			return event
		}
		return event
	})
}

func (c *CardPricesMode) SetupView(prices *ygoprices.CardCollection) {
	c.CardIndex = 0
	if c.Display.App == nil {
		fmt.Println("Display nil")
		c.Display = displaymanager.NewDisplayManager()
	}
	c.SetupInputCapture(len(prices.Cards), prices)
	display.DisplayCardQueryData(c.Display.App, c.Display.Flex, len(prices.Cards), c.CardIndex, prices.Cards[c.CardIndex])

	if err := c.Display.App.SetRoot(c.Display.Flex, true).SetFocus(c.Display.Flex).Run(); err != nil {
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
		c.Display = displaymanager.NewDisplayManager()
	} else if response == "Read" {
		c.NewPrice = false
		c.setCMode(Read)
		c.NewPrice = true
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
	stuff, err := cardStruct.ReadDataToSlice()
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

		nameToSearch, prices, _, err := SelectCardQuery()
		c.CardName = nameToSearch
		fmt.Println(nameToSearch)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(nameToSearch)

		pricesData := formatDataForOutput(prices)
		fmt.Printf("Returned: %d\n", len(pricesData))

		//Begin View Logic.
		c.Collection = prices

		c.SetupView(c.Collection)
		if err = c.Display.App.Run(); err != nil {
			fmt.Println(err)
			return
		}
		c.Display.App.Stop()
		c.Display.Flex.Clear()
		fmt.Println("Hello world!")
		fmt.Println(c.CardName)

		d := ygoprodeck.YuGiOhProDeckSearchData{
			Name:    c.CardName,
			CardSet: c.SetName,
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

func (c *CardPricesMode) Insert() (bool, error) {
	db, err := sql.Open("sqlite3", "./dist/card_data.db")
	if err != nil {
		return false, err
	}
	sqliteInsertStatement := `INSERT INTO card_data (CardName, CardSetName, PrintTag, CardPrice, High, Low, Average, Shift, Shift3, Shift7, Shift21, Shift30, Shift90, Shift180, Shift365, TimeLastUpdated, ImageUrl, CardURL, TrackedTime) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	statement, err := db.Prepare(sqliteInsertStatement)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
