package iphone

import (
	"fmt"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
)

type IIphone interface {
	Assemble()
	Certificates()
	Pack()
	GetHardware()
}

type AbstractIPhone struct {
	IIphone
	rules abstractfactory.ICountryRulesAbstractFactory
}

func (a AbstractIPhone) Assemble() {
	fmt.Println("Assembling all the hardwares")
}

func (a AbstractIPhone) Certificates() {
	fmt.Println("Testing all the certificates")
	fmt.Println(a.rules.GetCertificates().ApplyCertificate())
}

func (a AbstractIPhone) Pack() {
	fmt.Println("Packing the device")
	fmt.Println(a.rules.GetPacking().Pack())
}
