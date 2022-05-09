package factory

import (
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/iphone"
)

type IIPhoneFactory interface {
	OrderIPhone(level string, createIPhone CreateIPhone) iphone.IIphone
	FuncCreateIPhone() func(string) iphone.IIphone
}

type AbstractIPhoneFactory struct {
	IIPhoneFactory
	Rules abstractfactory.ICountryRulesAbstractFactory
}

type CreateIPhone func(string) iphone.IIphone

func (a AbstractIPhoneFactory) OrderIPhone(level string, createIPhone CreateIPhone) iphone.IIphone {

	device := createIPhone(level)

	//device.GetHardware()
	//device.Assemble()
	device.Certificates()
	device.Pack()

	return device
}
