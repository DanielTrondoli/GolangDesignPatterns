package iphone

import (
	"fmt"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
)

type IPhone11Pro struct {
	AbstractIPhone
}

func NewIPhone11Pro(rules abstractfactory.ICountryRulesAbstractFactory) *IPhone11Pro {
	new := new(IPhone11Pro)
	new.rules = rules
	return new
}

func (i IPhone11Pro) GetHardware() {
	fmt.Println("Hardware list")
	fmt.Println("\t- 6.1in Screen")
	fmt.Println("\t- A13 Chipset")
	fmt.Println("\t- 4Gb RAM")
	fmt.Println("\t- 256Gb Memory")
}
