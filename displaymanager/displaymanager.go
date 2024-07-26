package displaymanager

import "github.com/rivo/tview"

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
	App         *tview.Application
	Flex        *tview.Flex
	mode        DisplayMode
	displayData DisplayDataInterface
}

func (d *DisplayManager) InitManager() {
	d.App = tview.NewApplication()
	d.Flex = tview.NewFlex().SetDirection(tview.FlexRow)
}

func (d *DisplayManager) SetMode(mode DisplayMode) {
	d.mode = mode
}
