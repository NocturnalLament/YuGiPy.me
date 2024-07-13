package ygoprodeck

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Iilun/survey/v2"
	"github.com/NocturnalLament/yugigo/display"
	"github.com/rivo/tview"
)

// https://db.ygoprodeck.com/api/v7/cardinfo.php
var _ display.CardDataDisplay

type CardData struct {
	Data []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		FrameType   string `json:"frameType"`
		Description string `json:"desc"`
		ATK         int    `json:"atk"`
		DEF         int    `json:"def"`
		Level       int    `json:"level"`
		Race        string `json:"race"`
		Attribute   string `json:"attribute"`
		CardSets    []struct {
			SetName       string  `json:"set_name"`
			SetCode       string  `json:"set_code"`
			SetRarity     string  `json:"set_rarity"`
			SetRarityCode string  `json:"set_rarity_code"`
			SetPrice      float64 `json:"set_price,string"`
		} `json:"card_sets"`
		CardImages []struct {
			ID              int    `json:"id"`
			ImageURL        string `json:"image_url"`
			ImageURLSmall   string `json:"image_url_small"`
			ImageURLCropped string `json:"image_url_cropped"`
		} `json:"card_images"`
		CardPrices []struct {
			CardmarketPrice   float64 `json:"cardmarket_price,string"`
			TcgplayerPrice    float64 `json:"tcgplayer_price,string"`
			EbayPrice         float64 `json:"ebay_price,string"`
			AmazonPrice       float64 `json:"amazon_price,string"`
			CoolstuffincPrice float64 `json:"coolstuffinc_price,string"`
		} `json:"card_prices"`
	} `json:"data"`
}

func (c CardData) GetCardNames() []string {
	names := []string{}
	for _, card := range c.Data {
		names = append(names, card.Name)
	}
	return names
}

type YGoProDeckPrompts map[string]string

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
	LinkMarkers string
	Scale       int
	CardSet     string
	Archetype   string
	Banlist     string
	Sort        string
	Format      string
	Misc        YuGiOhProDeckSearchMisc   // Will either be unpassed or if true will be passed as "yes"
	Staple      YuGiOhProDeckSearchStaple // Will either be unpassed or if true will be passed as "yes"
}

func (y CardData) DisplayData() {

}

type YuGiOhProDeckSearchMisc bool

func (y YuGiOhProDeckSearchMisc) String() string {
	if y {
		return "yes"
	}
	return "no"
}

func (y YuGiOhProDeckSearchMisc) MarshalJSON() ([]byte, error) {
	if y {
		return []byte(`"yes"`), nil
	} else {
		return nil, fmt.Errorf("value is not 'yes', marshalling to JSON is not allowed")
	}
}

type YuGiOhProDeckSearchStaple bool

type CardArchetype string

func GetFilterPromptString(inputType string) string {
	return fmt.Sprintf("Enter the %s of the card", inputType)
}

func GetAtkPrompt() string {
	prompt := survey.Input{
		Message: GetFilterPromptString("ATK"),
	}
	atkPoints := ""
	err := survey.AskOne(&prompt, &atkPoints)
	if err != nil {
		fmt.Println(err.Error())
	}
	return atkPoints
}

func GetDefPrompt() string {
	prompt := survey.Input{
		Message: GetFilterPromptString("DEF"),
	}
	defPoints := ""
	err := survey.AskOne(&prompt, &defPoints)
	if err != nil {
		fmt.Println(err.Error())
	}
	return defPoints
}

func GetPendulumScalePrompt() string {
	scale := ""
	prompt := survey.Question{
		Name: "scale",
		Prompt: &survey.Input{
			Message: GetFilterPromptString("Scale"),
		},
		Validate: func(input interface{}) error {
			strInput, ok := input.(string)
			if !ok {
				return fmt.Errorf("invalid input")
			}
			scale, err := strconv.Atoi(strInput)
			if err != nil {
				return fmt.Errorf("invalid input")
			}

			if scale < 1 || scale > 8 {
				return fmt.Errorf("invalid input")
			}
			return nil
		},
	}

	err := survey.Ask([]*survey.Question{&prompt}, &scale)
	if err != nil {
		fmt.Println(err.Error())
	}
	return scale
}

func GetValsFromPrompt(selectedItems []string) YGoProDeckPrompts {
	response := make(map[string]string)
	for _, item := range selectedItems {
		switch item {
		case "Name":
			response["Name"] = GetFilterName()
		case "Fuzzy Name":
			response["fname"] = GetFuzzyNameFilter()
		case "ID":
			response["Id"] = GetCardIDPrompt()
		case "Konami ID":
			response["KonamiId"] = GetCardKonamiId()
		case "Type":
			response["Type"] = GetCardType()
		case "ATK":
			response["Atk"] = GetAtkPrompt()
		case "DEF":
			response["Def"] = GetDefPrompt()
		case "Level":
			response["Level"] = GetCardLevelPrompt()
		case "Attributes":
			response["Attributes"] = GetCardAttributes()
		case "Link":
			response["Link"] = GetLinkValuePrompt()
		case "LinkMarkers":
			response["LinkMarkers"] = GetLinkMarkers()
		case "Scale":
			response["Scale"] = GetPendulumScalePrompt()
		case "Card Set":
			response["CardSet"] = GetCardSetPrompt()
		case "Archetype":
			response["Archetype"] = PromptCardArchetype()
		case "Banlist":
			response["Banlist"] = GetBanList()
		case "Sort":
			response["sort"] = PromptSortBy()
		case "Format":
			response["Format"] = GetCardFormat()
		case "Misc":
			response["Misc"] = "yes"
		case "Staple":
			response["Staple"] = "yes"
		case "has_effect":
			response["has_effect"] = "yes"
		}
	}
	return response
}

func GetCardSetPrompt() string {
	prompt := survey.Input{
		Message: GetFilterPromptString("Card Set"),
	}
	cardSet := ""
	err := survey.AskOne(&prompt, &cardSet)
	if err != nil {
		fmt.Println(err.Error())
	}
	return cardSet
}

func ProDeckPrompt() []string {
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
	return selected
	//return nil, nil
}

func PromptSortBy() string {
	sortByOptions := []string{"Name", "ATK", "DEF", "Type", "Level", "Id", "New"}
	prompt := survey.Select{
		Message: "Select the sort by option",
		Options: sortByOptions,
	}
	selected := ""
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return selected
}

func (p *YGoProDeckPrompts) ProcessPrompts() *YuGiOhProDeckSearchData {
	ygoPro := NewYGOPRoDeckSearchData()
	for key, prompt := range *p {
		promptLower := strings.ToLower(key)

		vals := reflect.ValueOf(ygoPro).Elem()
		for i := 0; i < vals.NumField(); i++ {
			field := vals.Type().Field(i)
			//fmt.Println("Field Name: ", fieldName)
			fmt.Println("Prompt: ", prompt)
			fieldName := strings.ToLower(field.Name)
			if fieldName == promptLower {

				switch field.Type.Kind() {
				case reflect.String:
					fmt.Println("Prompt: ", prompt)
					vals.Field(i).SetString(prompt)
				case reflect.Int:
					if outVal, err := strconv.Atoi(prompt); err == nil {
						vals.Field(i).SetInt(int64(outVal))
					}
				case reflect.Bool:
					vals.Field(i).SetBool(true)
				}
			}
		}
	}
	return ygoPro
}

func (p YuGiOhProDeckSearchData) Mapify() map[string]string {
	result := make(map[string]string)
	vals := reflect.ValueOf(p)

	for i := 0; i < vals.NumField(); i++ {
		field := vals.Field(i)

		fieldName := vals.Type().Field(i).Name
		if misc, ok := field.Interface().(YuGiOhProDeckSearchMisc); ok {
			//result[fieldName] = "yes"
			//fieldValue := field.Interface().(bool)
			if misc {
				result[fieldName] = "yes"
			}
		} else if staple, ok := field.Interface().(YuGiOhProDeckSearchStaple); ok {
			//[fieldName] = "yes"
			if staple {
				result[fieldName] = "yes"
			}
		}
		switch field.Kind() {
		case reflect.String:
			if field.String() == "Default" {
				continue
			}
			result[strings.ToLower(fieldName)] = field.String()
		case reflect.Int:
			if field.Int() == 0 {
				continue
			}
			result[fieldName] = strconv.Itoa(int(field.Int()))
		}

	}

	return result
}

func (c CardData) DisplayCard(app *tview.Application, flex *tview.Flex, displayIndex int) {
	if len(c.Data) > 0 {
		card := c.Data[displayIndex]
		var output string
		output += fmt.Sprintf("Name: %s\n", card.Name)
		output += fmt.Sprintf("Type: %s\n", card.Type)
		output += fmt.Sprintf("ATK: %d\n", card.ATK)
		output += fmt.Sprintf("DEF: %d\n", card.DEF)
		output += fmt.Sprintf("Level: %d\n", card.Level)
		for _, set := range card.CardSets {
			output += fmt.Sprintf("Set Name: %s\n", set.SetName)
			output += fmt.Sprintf("Set Code: %s\n", set.SetCode)
			output += fmt.Sprintf("Set Rarity: %s\n", set.SetRarity)
			output += fmt.Sprintf("Set Price: %f\n", set.SetPrice)
		}
		textView := tview.NewTextView()
		textView.SetText(output)
		textView.SetBorder(true)
		textView.SetTitle("Card Information")
		textView.SetTitleAlign(tview.AlignCenter)
		textView.SetBorderPadding(1, 1, 2, 2)
		textView.SetDynamicColors(true)
		textView.SetWordWrap(true)
		textView.SetWrap(true)
		flex.AddItem(textView, 0, 1, false)
	}
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
