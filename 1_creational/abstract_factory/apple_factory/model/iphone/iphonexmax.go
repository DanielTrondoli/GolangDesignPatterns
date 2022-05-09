package iphone

import (
	"fmt"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
)

type IPhoneXMax struct {
	AbstractIPhone
}

func NewIphoneXMax(rules abstractfactory.ICountryRulesAbstractFactory) *IPhoneXMax {
	new := new(IPhoneXMax)
	new.rules = rules
	return new
}

func (i IPhoneXMax) GetHardware() {
	fmt.Println("Hardware list")
	fmt.Println("\t- 6.1in Screen")
	fmt.Println("\t- A13 Chipset")
	fmt.Println("\t- 4Gb RAM")
	fmt.Println("\t- 256Gb Memory")
}
