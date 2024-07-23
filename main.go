package main

import (
	"fmt"
	"github.com/NocturnalLament/yugigo/maininterface"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Get data to search
	//mode := PickMode()
	//m := ModeSwitch(mode)
	//m.Execute()
	p := maininterface.ProgramLayout{}
	p.InitMode()
	outThing := maininterface.CPricesMode.CardName
	fmt.Println(outThing)
	_, err := maininterface.CPricesMode.Insert()
	if err != nil {
		fmt.Println(err)
	}

}
