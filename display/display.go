package display

import (
	"github.com/rivo/tview"
)

type CardDataDisplay interface {
	DisplayData() string
}

//func DisplayData(app *tview.Application, flex *tview.Flex, data CardDataDisplay)

func ShowCardQueryData(flex *tview.Flex, length int, index int, data CardDataDisplay) {
	outData := data.DisplayData()
	text := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetScrollable(true).
		SetRegions(true)

	if index <= length-1 {
		flex.Clear()
		text.SetText(outData)
		flex.AddItem(text, 0, 1, false)
	}

}

func DisplayEndOfPrices(app *tview.Application, flex *tview.Flex) {
	flex.Clear()
	endMessage := "End of data. Press any key to return."
	text := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetScrollable(true).
		SetRegions(true)
	text.SetText(endMessage).SetTextAlign(tview.AlignCenter)
	flex.AddItem(text, 0, 1, false)
}

func DisplaySelectedCardPrice(app *tview.Application, flex *tview.Flex, data string) {
	textView := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetScrollable(true).
		SetRegions(true)
	textView.SetText(data)
	flex.Clear()
	flex.AddItem(textView, 0, 1, false)
}
