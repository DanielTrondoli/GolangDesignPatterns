package factory

import (
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/iphone"
)

type IPhoneXFactory struct {
	AbstractIPhoneFactory
}

func NewIPhoneXFactory(rules abstractfactory.ICountryRulesAbstractFactory) *IPhoneXFactory {
	new := new(IPhoneXFactory)
	new.Rules = rules
	return new
}

func (f IPhoneXFactory) FuncCreateIPhone() func(string) iphone.IIphone {
	return func(level string) iphone.IIphone {
		if level == "standard" {
			return iphone.NewIphoneX(f.Rules)
		} else if level == "max" {
			return iphone.NewIphoneXMax(f.Rules)
		}
		return nil
	}
}
