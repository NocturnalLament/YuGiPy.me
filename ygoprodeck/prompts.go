package ygoprodeck

import (
	"fmt"
	"strings"

	"github.com/Iilun/survey/v2"
	"github.com/davecgh/go-spew/spew"
)

func GetDataToSearch() (*YuGiOhProDeckSearchData, error) {

	items := ProDeckPrompt()
	vals := GetValsFromPrompt(items)
	item := vals.ProcessPrompts()
	if item == nil {
		return nil, fmt.Errorf("error processing prompts")
	}
	spew.Dump(item)
	fmt.Printf("Item: %v\n", item.Name)
	return item, nil
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
