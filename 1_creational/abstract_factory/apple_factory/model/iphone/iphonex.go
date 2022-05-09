package iphone

import (
	"fmt"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
)

type IPhoneX struct {
	AbstractIPhone
}

func NewIphoneX(rules abstractfactory.ICountryRulesAbstractFactory) *IPhoneX {
	new := new(IPhoneX)
	new.rules = rules
	return new
}

func (i IPhoneX) GetHardware() {
	fmt.Println("Hardware list")
	fmt.Println("\t- 6.1in Screen")
	fmt.Println("\t- A13 Chipset")
	fmt.Println("\t- 4Gb RAM")
	fmt.Println("\t- 256Gb Memory")
}
