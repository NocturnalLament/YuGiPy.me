package displaymanager

import (
	"github.com/NocturnalLament/yugigo/display"
	"github.com/rivo/tview"
)

type DisplayMode int

const (
	DefaultDisplay DisplayMode = iota
	CardData
	CardPrices
)

type DisplayDataInterface interface {
	Display()
}

type DisplayManager struct {
	App             *tview.Application
	Flex            *tview.Flex
	mode            DisplayMode
	displayData     *display.CardDataDisplay
	displayCallback func(dataInterface *DisplayDataInterface)
}

func (d *DisplayManager) InitManager() {
	d.App = tview.NewApplication()
	d.Flex = tview.NewFlex().SetDirection(tview.FlexRow)
}

func (d *DisplayManager) SetMode(mode DisplayMode) {
	d.mode = mode
}

func (d *DisplayManager) SetData(data *display.CardDataDisplay) {
	d.displayData = data
}

func (d *DisplayManager) SetCallback(callback func(dataInterface *DisplayDataInterface)) {
	d.displayCallback = callback
}

func NewDisplayManager() *DisplayManager {
	return &DisplayManager{
		App:             tview.NewApplication(),
		Flex:            tview.NewFlex().SetDirection(tview.FlexRow),
		mode:            DefaultDisplay,
		displayData:     nil,
		displayCallback: nil,
	}
}
