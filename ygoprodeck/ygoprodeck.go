package ygoprodeck

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Iilun/survey/v2"
)

// https://db.ygoprodeck.com/api/v7/cardinfo.php

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

func NewYGOPRoDeckSearchData() *YuGiOhProDeckSearchData {
	return &YuGiOhProDeckSearchData{
		Name:        "Default",
		FName:       "Default",
		Id:          0,
		KonamiId:    0,
		Type:        "Default",
		Atk:         0,
		Def:         0,
		Level:       0,
		Race:        "Default",
		Attribute:   "Default",
		Link:        "Default",
		LinkMarkers: "Default",
		Scale:       0,
		CardSet:     "Default",
		Archetype:   "Default",
		Banlist:     "Default",
		Sort:        "Default",
		Format:      "Default",
		Misc:        false,
		Staple:      false,
	}
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
	selected := []string{}
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	selectedAttributes := strings.Join(selected, ",")
	return selectedAttributes
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

func GetCardFormat() string {
	cardFormats := []string{"OCG", "TCG", "Speed Duel", "Rush Duel", "Goat", "OCG Goat", "Duel Links"}
	prompt := survey.Select{
		Message: "Select the card format",
		Options: cardFormats,
	}
	selected := ""
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	return selected
}

type CardArchetype string

func PromptCardArchetype() string {
	cardArchetypes := []string{
		"Albaz", "Ally of Justice", "Appliancer", "Ashened", "Abyss Actor", "Altergeist",
		"Aquaactress", "Assault Mode", "Adamancipator", "Amazement", "Arcana Force", "Atlantean",
		"Adventurer Token", "Amazoness", "Archfiend", "Aesir", "Amorphage", "Armed Dragon", "Agent", "Ancient Gear", "Aroma",
		"Alien", "Ancient Warriors", "Artifact", "Batteryman", "Battlewasp", "Battlin Boxer", "Beetrooper", "Black Luster", "Blackwing",
		"Blue-Eyes", "Bounzer", "Bujin", "Burning Abyss", "Buster Blader", "Butterspy", "Bystial", "Centur-Ion", "Chaos", "Chemicritter",
		"Chronomaly", "Chrysalis", "Cipher", "Cloudian", "Code Talker", "Constellar", "Crusadia", "Crystal Beast", "Crystron", "Cubic", "Cyberdark",
		"Cyber Dragon", "D/D/", "Danger!", "Darklord", "Dark Magician", "Dark Scorpion", "Dark World", "D.D.", "Deep Sea", "Deskbot", "Despia", "Destiny Hero",
		"Digital Bug", "Dinomist", "Dinomorphia", "Dinowrestler", "Dododo", "Dogmatika", "Doodle Beast", "Dracoslayer", "Dragon Ruler", "Dragonmaid", "Dragunity",
		"Dream Mirror", "Dryton", "Dual Avatar", "Duston", "Earthbound", "Edge Imp", "Edlich", "Elemental Hero", "Elementsaber", "Empowered Warrior", "Endymion",
		"Evil Eye", "Evil Hero", "Evil Twin / Live Twin", "Evilswarm", "Evoltile", "Exosister", "Eyes Restrict", "Fabled", "Face Cards", "F.A.", "Fairy Tail", "Familiar-Posessed",
		"Brotherhood of the Fire Fist", "Fire King", "Fire Warrior", "Flame Swordsman", "Flamvell", "Fleur", "Floowandereeze", "Flower Cardian", "Fluffal", "Forbidden One (Exodia)",
		"Fortune Fairy", "Fortune Lady", "Fossil Fusion", "Frightfur", "Frog", "Fur Hire", "G Golem", "Gadget", "Gagaga", "Gaia", "Galaxy", "Ganbara", "Gate Guardian", "Gearfried", "Geargia",
		"Gem", "Generaider", "Ghostrick", "Ghoti", "Gimmick Puppet", "Gishki", "Glacial Beast", "Gladiator Beast", "Goblin", "Goblin Biker", "Gogogo", "Gold Pride", "Gorgonic", "Gouki", "Goyo", "Gravekeeper",
		"Graydle", "Gunkan", "Gusto", "Harpie", "Hazy Flame", "Heraldic Beast", "Heroic", "Hieratic", "Horus", "Ice Barrier", "Icejade", "Ignknight", "Ignister", "Impcantation", "Infernity", "Infernoble", "Infernoid",
		"Infinitrack", "Invoked", "Inzektor", "Iron Chain", "Junk", "Jurrac", "Kaiju", "Karakuri", "Kashtira", "Knightmare", "Koa'ki Meiru", "Kozmo", "Krawler", "Krawler", "Kuriboh", "Labrynth", "Laval", "Libromancer", "Lightray",
		"Lightsworn", "Lswarm", "Lunalight", "Lyrilusc", "Machina", "Madolche", "Magical Musket", "Magician", "Magikey", "Magistus", "Majespecter", "Malefic", "Mannadium", "Marincess", "Masked HERO", "Materiactor", "Mathmech", "Mayakashi",
		"Mekk-Knight", "Meklord", "Melffy", "Melodious", "Memento", "Mermail", "Metalfoes", "Metaphys", "Mikanko", "Mist Valley", "Monarch", "Morphtronic", "Mystical Beast", "Mythical Beast", "Myutant", "Naturia", "Nekroz", "Nemleria", "Neo-Spacian",
		"Neos", "Nemeses", "Nephthys", "Nimble", "Ninja", "Noble Knight", "Nordic", "Nouvelles", "Numeron", "Number", "Odd-Eyes", "Ogdoadic", "Orcust", "Ojama", "P.U.N.K.", "Paleozoic", "Parshath", "Penguin", "Performage", "Performapal", "Phantasm Spiral",
		"Phantom Beast", "Photon", "Plunder Patroll", "Prank-Kids", "Predaplant", "Prediction Princess", "Purrely", "PSY-Framegear", "Psychic", "Qli", "Ragnaraika", "Raidraptor", "Red-Eyes", "Reptilianne", "Rescue-ACE", "Resonator", "Rikka", "Risebell",
		"Ritual Beast", "Roid", "Rokket", "Rose", "Runick", "Sangen", "S-Force", "Salamangreat", "Scareclaw", "Scrap", "Shaddoll", "Shark", "Shining Sarcophagus", "Shinobird", "Shiranui", "Silent Magician", "Silent Swordsman", "Simorgh", "Sinful Spoils",
		"Six Samurai", "Skull Servant", "Sky Striker", "Snake-Eye", "Solfachord", "Speedroid", "Spellbook", "Springan", "Spyral", "Spellcaster", "Spright", "Star Sereph", "Starry Knight", "Steelswarm", "Subterror", "Sunavalon", "Superheavy Samurai",
		"Supreme King", "Swordsoul", "Sylvan", "Symphonic Warrior", "Synchron", "Tearlamants", "Tellarknight", "Tenyi", "T.G.", "The Agent", "The Phantom Knights", "The Weather", "Therion", "Thunder Dragon", "Time Thief", "Timelord", "Tindangle", "Tistina",
		"Toon", "Traptrix", "Triamid", "Tri-Brigade", "Trickstar", "True King", "Twilightsworn", "U.A.", "Unchained", "Ursarctic", "Utopia", "Vaalmonica", "Vampire", "Vanquish Soul", "Vassal", "Vaylantz", "Vendread", "Venom", "Virtual World", "Visas Starfrost",
		"Vision HERO", "Voiceless Voice", "Volcanic", "Vylon", "War Rock", "Watt", "Wind-up", "Windwitch", "Witchcrafter", "World Chalice", "World Legacy", "Worm", "Xtra HERO", "Xyz", "X-Saber", "Yang Zing", "Yosenju", "Yubel", "Zefra", "Zoodiac", "ZW", "Zombie",
		"Zubaba",
	}
	prompt := survey.MultiSelect{
		Message:  "Select the card archetype",
		Options:  cardArchetypes,
		PageSize: 10,
	}
	selected := []string{}
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	joinedSelected := strings.Join(selected, ",")
	return joinedSelected
}

func GetFilterName() string {
	prompt := survey.Input{
		Message: "Enter the name of the card",
	}
	userInput := ""
	err := survey.AskOne(&prompt, &userInput)
	if err != nil {
		fmt.Println(err.Error())
	}
	return userInput
}

func GetCardLevelPrompt() string {
	prompt := survey.Input{
		Message: GetFilterPromptString("Level"),
	}
	input := ""
	err := survey.AskOne(&prompt, &input)
	if err != nil {
		fmt.Println(err.Error())
	}
	return input
}

func GetFuzzyNameFilter() string {
	prompt := survey.Input{
		Message: GetFilterPromptString("Fuzzy Name"),
	}
	selected := ""
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	return selected
}

func GetCardIDPrompt() string {
	prompt := survey.Input{
		Message: GetFilterPromptString("ID"),
	}
	selected := ""
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	return selected
}

func GetCardKonamiId() string {
	prompt := survey.Input{
		Message: GetFilterPromptString("Konami ID"),
	}
	selected := ""
	err := survey.AskOne(&prompt, &selected)
	if err != nil {
		fmt.Println(err.Error())
	}
	return selected
}

func GetCardType() string {
	cardTypes := []string{"Normal Monster", "Effect Monster", "Synchro Monster", "Spell Card", "Trap Card", "Flip Tuner Monster", "Flip Tuner Effect Monster", "Gemini Monster",
		"Normal Tuner Monster", "Pedulum Effect Monster", "Pendulum Effect Ritual Monster", "Pendulum Flip Effect Monster", "Pendulum Normal Monster", "Pendulum Tuner Effect Monster",
		"Ritual Effect Monster", "Ritual Monster", "Spirit Monster", "Toon Monster", "Tuner Monster", "Union Effect Monster", "Fusion Monster", "Link Monster", "Pendulum Effect Fusion Monster",
		"Synchro Monster", "Synchro Pendulum Effect Monster", "XYZ Monster", "XYZ Pendulum Effect Monster", "Skill Card", "Token"}
	selectedCards := []string{}
	prompt := survey.MultiSelect{
		Message:  GetFilterPromptString("type"),
		Options:  cardTypes,
		PageSize: 10,
	}
	err := survey.AskOne(&prompt, &selectedCards)
	if err != nil {
		fmt.Println(err.Error())
	}
	outCards := strings.Join(selectedCards, ",")
	return outCards
}

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

func GetLinkValuePrompt() string {
	prompt := survey.Input{
		Message: GetFilterPromptString("Link Value"),
	}

	linkValue := ""
	err := survey.AskOne(&prompt, &linkValue)
	if err != nil {
		fmt.Println(err.Error())
	}
	return linkValue
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
