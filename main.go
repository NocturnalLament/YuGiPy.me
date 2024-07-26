package main

import (
	"github.com/NocturnalLament/yugigo/maininterface"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Get data to search
	//mode := PickMode()
	//m := ModeSwitch(mode)
	//m.Execute()
	p := maininterface.ProgramLayout{}
	modeStr := maininterface.PickMode()
	modeConst := maininterface.GetExecConstant(modeStr)
	p.ExecutionConst = modeConst
	p.InitMode()

}
