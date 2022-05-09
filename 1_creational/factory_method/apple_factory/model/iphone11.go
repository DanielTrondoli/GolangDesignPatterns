package model

import (
	"fmt"
)

type IPhone11 struct {
	AbstracteIPhone
}

func NewIPhone11(generation, level string) *IPhone11 {
	new := new(IPhone11)
	new.InputModel(generation, level)
	return new
}

func (i IPhone11) GetHardware() {
	fmt.Println("Hardware list")
	fmt.Println("\t- 6.1 in Screen")
	fmt.Println("\t- A13 Chipset")
	fmt.Println("\t- 4Gb RAM")
	fmt.Println("\t- 256Gb Memory")
}
