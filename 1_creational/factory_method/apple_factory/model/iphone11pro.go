package model

import (
	"fmt"
)

type IPhone11Pro struct {
	AbstracteIPhone
}

func NewIPhone11Pro(generation, level string) *IPhone11Pro {
	new := new(IPhone11Pro)
	new.InputModel(generation, level)
	return new
}

func (i IPhone11Pro) GetHardware() {
	fmt.Println("Hardware list")
	fmt.Println("\t- 6.1 in Screen")
	fmt.Println("\t- A13 Chipset")
	fmt.Println("\t- 4Gb RAM")
	fmt.Println("\t- 256Gb Memory")
}
