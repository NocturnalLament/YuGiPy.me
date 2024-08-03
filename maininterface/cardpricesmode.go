package maininterface

import (
	"database/sql"
	"fmt"
	survey "github.com/Iilun/survey/v2"
	"github.com/NocturnalLament/yugigo/display"
	"github.com/NocturnalLament/yugigo/displaymanager"
	"github.com/NocturnalLament/yugigo/timelogic"
	"github.com/NocturnalLament/yugigo/writemanager"
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
	InsertionPrepared    bool
	CardData             *ygoprices.Card
	Display              *displaymanager.DisplayManager
	DisplaySetupCallback func()
	SqliteDataRecord     *ygoprices.YgoPricesCardData
	cardSelected         bool
	CardUrl              string
	NewPrice             bool
	CardSelected         bool
	Collection           *ygoprices.CardCollection
	LoadData             *ygoprices.YgoPricesCardData
	WriteManager         *writemanager.WriteManager
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
			fmt.Println("SqliteDataRecord read successfully")
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
		SqliteDataRecord: nil,
		cardSelected:     false,
		CardUrl:          "",
		NewPrice:         false,
		WriteManager:     writemanager.NewManager(),
	}
}

func (c *CardPricesMode) SetupView(prices *ygoprices.CardCollection) {
	c.CardIndex = 0
	if c.Display.App == nil {
		fmt.Println("Display nil")
		c.Display = displaymanager.NewDisplayManager()
	}
	c.SetupInputCapture(len(prices.Cards), prices)
	display.ShowCardQueryData(c.Display.Flex, len(prices.Cards), c.CardIndex, prices.Cards[c.CardIndex])

	if err := c.Display.App.SetRoot(c.Display.Flex, true).SetFocus(c.Display.Flex).Run(); err != nil {
		panic(err)
	}
	c.returnToConsole()
}

func (c *CardPricesMode) initializeMode() {
	options := []string{"Read", "Write", "Load"}
	prompt := survey.Select{
		Message: "Read/Write/Load SqliteDataRecord?",
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
	} else if response == "Load" {
		return
	}
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
			/* if c.CardSelected == true {
				_, err := c.Insert()
				if err != nil {
					fmt.Println(err)
				}
				c.CardSelected
			} */
			if c.InsertionPrepared && c.cardSelected {
				err := c.Write()
				if err != nil {
					panic(err)
				}
				c.returnToConsole()
			} else if c.cardSelected {
				c.Display.Flex.Clear()
				display.DisplayEndOfPrices(c.Display.App, c.Display.Flex)

				c.InsertionPrepared = true
			} else if c.CardIndex < amountOfCards {
				c.Display.Flex.Clear()
				display.ShowCardQueryData(c.Display.Flex, len(prices.Cards), c.CardIndex, prices.Cards[c.CardIndex])
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
			s := ygoprodeck.YuGiOhProDeckSearchData{
				Name:    c.CardName,
				CardSet: c.SetName,
			}
			f := ygoprodeck.URLAttrBuilder(&s)
			q, err := ygoprodeck.Query(f)
			if err != nil {
				fmt.Println(err)
				return event
			}
			fmt.Println(len(q.Data))
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
				CardURL:         c.CardUrl,
				TimeLastUpdated: timelogic.GetNewTime(),
			}
			c.SqliteDataRecord = &y

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

func (c *CardPricesMode) returnToConsole() {
	// Stop the tview application
	if c.Display != nil && c.Display.App != nil {
		c.Display.App.Stop()
	}

	// Clear the tview Flex layout
	if c.Display != nil && c.Display.Flex != nil {
		c.Display.Flex.Clear()
	}

	// Print a message to indicate returning to the normal console
	fmt.Println("Returning to the normal console...")
}

func (c *CardPricesMode) ReadData() (bool, error) {
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

func (c *CardPricesMode) Execute() {
	c.initializeMode()
	c.executionModeSwitch()
}

func (c *CardPricesMode) executionModeSwitch() {
	if c.NewPrice {

		err, done := c.QueryData()
		if err != nil {
			fmt.Println(err)
			return
		}
		if done {
			return
		}
		c.SetupView(c.Collection)
	} else {
		fmt.Println("Exiting...")
	}
}

func (c *CardPricesMode) QueryData() (error, bool) {
	nameToSearch, prices, err, done := c.GetCardQueryData()
	if done {
		return nil, true
	}
	c.CardName = nameToSearch
	//Begin View Logic.
	c.Collection = prices
	return err, false
}

func (c *CardPricesMode) GetCardQueryData() (string, *ygoprices.CardCollection, error, bool) {
	nameToSearch, prices, _, err := SelectCardQuery()
	if err != nil {
		fmt.Println(err)
		return "", nil, nil, true
	}
	return nameToSearch, prices, err, false
}

func (c *CardPricesMode) Write() error {
	db, err := sql.Open("sqlite3", "./dist/card_data.db")
	if err != nil {
		return err
	}
	sqliteInsertStatement := `INSERT INTO card_data (CardName, CardSetName, PrintTag, CardPrice, High, Low, Average, Shift, Shift3, Shift7, Shift21, Shift30, Shift90, Shift180, Shift365, TimeLastUpdated, ImageUrl, CardURL, TrackedTime) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	statement, err := db.Prepare(sqliteInsertStatement)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(statement)
	_, err = statement.Exec(c.CardName, c.SqliteDataRecord.CardName, c.SqliteDataRecord.PrintTag, c.SqliteDataRecord.CardPrice, c.SqliteDataRecord.High, c.SqliteDataRecord.Low,
		c.SqliteDataRecord.Average, c.SqliteDataRecord.Shift, c.SqliteDataRecord.Shift3, c.SqliteDataRecord.Shift7, c.SqliteDataRecord.Shift21, c.SqliteDataRecord.Shift30, c.SqliteDataRecord.Shift90,
		c.SqliteDataRecord.Shift180, c.SqliteDataRecord.Shift365, c.SqliteDataRecord.TimeLastUpdated, c.SqliteDataRecord.ImageUrl, c.CardUrl, c.SqliteDataRecord.TrackedTime)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
