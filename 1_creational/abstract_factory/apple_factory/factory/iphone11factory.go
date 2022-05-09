package factory

import (
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/iphone"
)

type IPhone11Factory struct {
	AbstractIPhoneFactory
	Create CreateIPhone
}

func NewIPhone11Factory(rules abstractfactory.ICountryRulesAbstractFactory) *IPhone11Factory {
	new := new(IPhone11Factory)
	new.Rules = rules
	return new
}

func (f IPhone11Factory) FuncCreateIphone() func(string) iphone.IIphone {
	return func(level string) iphone.IIphone {
		if level == "standard" {
			return iphone.NewIphone11(f.Rules)
		} else if level == "pro" {
			return iphone.NewIPhone11Pro(f.Rules)
		}
		return nil
	}
}
