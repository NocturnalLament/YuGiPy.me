package maininterface

import (
	"database/sql"
	"fmt"
	"github.com/Iilun/survey/v2"
	"github.com/NocturnalLament/yugigo/displaymanager"

	"github.com/NocturnalLament/yugigo/ygoprices"
	"github.com/NocturnalLament/yugigo/ygoprodeck"
)

var ExecMode ExecutionMode

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
	SearchData           *ygoprodeck.YuGiOhProDeckSearchData
	ReturnedCardData     *ygoprodeck.CardData
	Display              *displaymanager.DisplayManager
	DisplaySetupCallback func()
	CardSelected         bool
	CurrentCardIndex     int
}

func (c *CardDataMode) ModeSwitch() {
	//TODO implement me
	panic("implement me")
}

func (c *CardDataMode) InitMode() {
	//TODO implement me
	panic("implement me")
}

type SubmodeOperator int

const (
	Default SubmodeOperator = iota
	Insert
	Read
	Update
)

type PriceMode struct {
	Mode          SubmodeOperator
	PriceData     *ygoprices.YgoPricesCardData
	DataInsertion func()
}

func (p *PriceMode) InitMode() {
	//TODO implement me
	fmt.Println("hi")
	items := []string{"Read", "Write"}
	prompt := survey.Select{
		Message: "Select a mode to run in:",
		Options: items,
	}
	var result string
	if err := survey.AskOne(&prompt, &result); err != nil {
		return
	}
	switch result {
	case "Read":
		p.Mode = Read
	case "Write":
		p.Mode = Insert
	}
	p.ModeSwitch()
}

func (p *PriceMode) ModeSwitch() {
	switch p.Mode {
	case Read:
		p.ReadData()
	case Insert:
		//TODO: Implement a way to get the CardPricesMode from
		c := CardPricesMode{}
		c.Execute()
	}
}

func (p *PriceMode) Execute() {
	p.InitMode()
}

func (p *PriceMode) UpdateMode(mode SubmodeOperator) {
	p.Mode = mode
}

func (p *PriceMode) ReadData() {
	_, err := p.PriceData.ReadData()
	if err != nil {
		return
	}
	return
}

type ProgramSubmode interface {
	ExecutionMode
	ModeSwitch()
	InitMode()
}

type ProgramModeDecision int

const (
	NoMode ProgramModeDecision = iota
	Price
	DataOperation
)

type Submode struct {
	modeOperator ProgramModeDecision
}

func (s *Submode) SetMode(p ProgramModeDecision) {
	s.modeOperator = p
}

func (s *Submode) ModeSwitch() {
	switch s.modeOperator {
	case Price:
		ExecMode = &PriceMode{}
	case DataOperation:
		ExecMode = &CardDataMode{}
	}
}

type DataOperations interface {
	ReadData()
	WriteData()
}

type PriceSubmode interface {
	ExecutionMode
	DataOperations
	ModeSwitch()
}

type PriceLoader interface {
	LoadSql(rows *sql.Rows) error
}

type CardTrackingData struct {
	CardName    string
	CardSetName string
	CardUrl     string
}

func (c *CardTrackingData) LoadSql(rows *sql.Rows) error {
	err := rows.Scan(&c.CardName, &c.CardSetName, &c.CardUrl)
	if err != nil {
		return err
	}
	return nil
}

func formatDataForOutput(prices *ygoprices.CardCollection) []*ygoprices.YgoPricesCardData {
	pricesData := []*ygoprices.YgoPricesCardData{}
	for _, card := range prices.Cards {
		priceDataStruct := ygoprices.NewYgoPriceData()
		priceDataStruct.CardName = card.Name

		priceDataStruct.PrintTag = card.PrintTag
		priceDataStruct.CardPrice = ygoprices.YGOCardPrice(card.PriceData.Data.Prices.Average)
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
	return pricesData
}

func PickMode() string {
	modes := []string{"Card Search", "Card Prices", "Server"}
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

type ExecConstant int

const (
	CardSearch ExecConstant = iota
	CardPrices
	Server
	None
)

func GetExecConstant(modeString string) ExecConstant {
	switch modeString {
	case "Card Search":
		return CardSearch
	case "Card Prices":
		return CardPrices
	case "Server":
		return Server
	}
	return None
}

type ProgramLayout struct {
	Mode           ExecutionMode
	SubmodeItem    Submode
	ExecutionConst ExecConstant
	Display        *displaymanager.DisplayManager
}

// Get Data in program layout
func (p *ProgramLayout) CreateDisplayInput() {
	switch mode := p.Mode.(type) {
	case *CardDataMode:
		mode.SetupInputCapture(mode.CurrentCardIndex, len(mode.ReturnedCardData.Data), mode.ReturnedCardData,
			p.Display.App, p.Display.Flex)
	case *CardPricesMode:
		mode.SetupInputCapture(len(mode.Prices), mode.Collection)
	}
}

func (p *ProgramLayout) ChangeExecConstant(e ExecConstant) {
	if e <= 3 && e >= 0 {
		p.ExecutionConst = e
	}
}

func (p *ProgramLayout) ChangeMode(m ExecutionMode) {
	p.Mode = m
	p.initSubmode()
	switch mode := p.Mode.(type) {
	case *CardPricesMode:
		p.Mode.Execute()
		mode.Display = displaymanager.NewDisplayManager()
	}
}

func (p *ProgramLayout) ModeSwitch() {
	switch p.ExecutionConst {
	case CardSearch:
		p.ChangeMode(NewCPricesMode())
	case CardPrices:

		p.ChangeMode(NewCPricesMode())

	}
}

func (p *ProgramLayout) AssignNewMode(modeString string) {
	switch modeString {
	case "Card Prices":
		p.ChangeMode(&PriceMode{})
	case "Card Search":
		p.ChangeExecConstant(CardSearch)
	}
}

func (p *ProgramLayout) InitMode() {
	//m := ModeSwitch(modeStr)

	p.ModeSwitch()
}

func (p *ProgramLayout) initSubmode() {

	switch mode := p.Mode.(type) {
	case *CardPricesMode:
		mode.DisplaySetupCallback = p.CreateDisplayInput
	case *CardDataMode:
		mode.DisplaySetupCallback = p.CreateDisplayInput
	}
}
