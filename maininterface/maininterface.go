package maininterface

import (
	"database/sql"
	"fmt"
	"github.com/Iilun/survey/v2"
	"github.com/NocturnalLament/yugigo/ygoprices"
	"github.com/NocturnalLament/yugigo/ygoprodeck"
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

type SubmodeOperator int

const (
	NoOperation SubmodeOperator = iota
	Insert
	Update
	Read
	Tracking
)

type PriceMode struct {
	Mode SubmodeOperator
}

func (p *PriceMode) UpdateMode(mode SubmodeOperator) {
	p.Mode = mode
}

func (p *PriceMode) ReadData() {
	return
}

type PriceSubmode interface {
	ExecutionMode
	ReadData()
	WriteData()
	ModeSwitch()
}

type DataType struct {
	SqlString string
}

type DataMode interface {
	Insert()
	Read()
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

func ModeSwitch(mode string) ExecutionMode {
	m := ExecutionMode(nil)
	switch mode {
	case "Card Search":
		m = NewCPricesMode()
	case "Card Prices":

		m = &CardPricesMode{}

	}
	return m
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

type ProgramLayout struct {
	Mode           ExecutionMode
	ExecutionConst ExecConstant
}

func (p *ProgramLayout) NewLayoutDelivery() *ProgramLayout {
	return &ProgramLayout{
		Mode:           nil,
		ExecutionConst: None,
	}
}

func (p *ProgramLayout) ChangeExecConstant(e ExecConstant) {
	if e <= 3 && e >= 0 {
		p.ExecutionConst = e
	}
}

func (p *ProgramLayout) ChangeMode(m ExecutionMode) {
	p.Mode = m
	m.Execute()
}

func (p ProgramLayout) InitMode() {
	modeStr := PickMode()
	m := ModeSwitch(modeStr)
	p.ChangeMode(m)
}
