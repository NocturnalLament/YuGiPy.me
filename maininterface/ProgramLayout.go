package maininterface

import "github.com/NocturnalLament/yugigo/displaymanager"

type ProgramLayout struct {
	Mode           ExecutionMode
	ExecutionConst ExecConstant
	Display        *displaymanager.DisplayManager
}

func (p *ProgramLayout) InitMode() {
	//m := ModeSwitch(modeStr)

	p.ModeSwitch()
}

func (p *ProgramLayout) ModeSwitch() {
	switch p.ExecutionConst {
	case CardSearch:
		p.ChangeMode(NewCDataMode())
	case CardPrices:

		p.ChangeMode(NewCPricesMode())

	}
}

func (p *ProgramLayout) ChangeMode(m ExecutionMode) {
	p.Mode = m
	p.initSubmode()
	switch mode := p.Mode.(type) {
	case *CardPricesMode:

		mode.Display = displaymanager.NewDisplayManager()
		p.Mode.Execute()
	case *CardDataMode:
		mode.Display = displaymanager.NewDisplayManager()
		mode.Execute()
	}
}

func (p *ProgramLayout) initSubmode() {

	switch mode := p.Mode.(type) {
	case *CardPricesMode:
		mode.DisplaySetupCallback = p.CreateDisplayInput
	case *CardDataMode:
		mode.DisplaySetupCallback = p.CreateDisplayInput
	}
}

// Get SqliteDataRecord in program layout
func (p *ProgramLayout) CreateDisplayInput() {
	//switch mode := p.Mode.(type) {
	//case *CardDataMode:
	//	mode.SetupInputCapture(mode.CurrentCardIndex, len(mode.ReturnedCardData.SqliteDataRecord), mode.ReturnedCardData,
	//		p.Display.App, p.Display.Flex)
	//case *CardPricesMode:
	//	mode.SetupInputCapture(len(mode.Prices), mode.Collection)
	//}
	return
}
