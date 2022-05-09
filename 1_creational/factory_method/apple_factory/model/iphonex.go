package model

import (
	"fmt"
)

type IPhoneX struct {
	AbstracteIPhone
}

func NewIPhoneX(generation, level string) *IPhoneX {
	new := new(IPhoneX)
	new.InputModel(generation, level)
	return new
}

func (i IPhoneX) GetHardware() {
	fmt.Println("Hardware list")
	fmt.Println("\t- 5.8in Screen")
	fmt.Println("\t- A11 Chipset")
	fmt.Println("\t- 3Gb RAM")
	fmt.Println("\t- 256Gb Memory")
}
