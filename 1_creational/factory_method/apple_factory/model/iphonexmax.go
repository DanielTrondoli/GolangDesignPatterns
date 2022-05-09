package model

import (
	"fmt"
)

type IPhoneXMax struct {
	AbstracteIPhone
}

func NewIPhoneXMax(generation, level string) *IPhoneXMax {
	new := new(IPhoneXMax)
	new.InputModel(generation, level)
	return new
}

func (i IPhoneXMax) GetHardware() {
	fmt.Println("Hardware list")
	fmt.Println("\t- 5.8in Screen")
	fmt.Println("\t- A11 Chipset")
	fmt.Println("\t- 3Gb RAM")
	fmt.Println("\t- 256Gb Memory")
}
