package display

import "github.com/rivo/tview"

type CardDataDisplay interface {
	DisplayData() string
}

//func DisplayData(app *tview.Application, flex *tview.Flex, data CardDataDisplay)

func DisplayCardQueryData(app *tview.Application, flex *tview.Flex, data CardDataDisplay) {
	outData := data.DisplayData()
	text := tview.NewTextView().
		SetText(outData).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetScrollable(true).
		SetRegions(true)
	flex.AddItem(text, 0, 1, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
