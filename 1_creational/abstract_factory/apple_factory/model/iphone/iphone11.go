package iphone

import (
	"fmt"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
)

type IPhone11 struct {
	AbstractIPhone
}

func NewIphone11(rules abstractfactory.ICountryRulesAbstractFactory) *IPhone11 {
	new := new(IPhone11)
	new.rules = rules
	return new
}

func (i IPhone11) GetHardware() {
	fmt.Println("Hardware list")
	fmt.Println("\t- 6.1in Screen")
	fmt.Println("\t- A13 Chipset")
	fmt.Println("\t- 4Gb RAM")
	fmt.Println("\t- 256Gb Memory")
}
