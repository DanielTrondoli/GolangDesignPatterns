package applefactory_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/factory/abstractfactory"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/apple_factory/model/iphone"
	"github.com/stretchr/testify/assert"
)

func TestAppleAbstractFactory(t *testing.T) {
	var rules abstractfactory.ICountryRulesAbstractFactory
	rules = new(abstractfactory.BrazilianRulesAbstractFactory)

	iphone11Factory := factory.NewIPhone11Factory(rules)
	iphoneXFactory := factory.NewIPhoneXFactory(rules)

	fmt.Println("### Ordering an iPhone 11")
	newIPhone11 := iphone11Factory.OrderIPhone("standard", iphone11Factory.FuncCreateIphone())
	newIPhone11Pro := iphone11Factory.OrderIPhone("pro", iphone11Factory.FuncCreateIphone())

	fmt.Println("### Ordering an iPhone X")
	newIPhoneX := iphoneXFactory.OrderIPhone("standard", iphoneXFactory.FuncCreateIPhone())
	newIPhoneXMax := iphoneXFactory.OrderIPhone("max", iphoneXFactory.FuncCreateIPhone())

	assert.Equal(t, reflect.TypeOf(newIPhoneX), reflect.TypeOf(iphone.NewIphoneX(rules)))
	assert.Equal(t, reflect.TypeOf(newIPhoneXMax), reflect.TypeOf(iphone.NewIphoneXMax(rules)))

	assert.Equal(t, reflect.TypeOf(newIPhone11), reflect.TypeOf(iphone.NewIphone11(rules)))
	assert.Equal(t, reflect.TypeOf(newIPhone11Pro), reflect.TypeOf(iphone.NewIPhone11Pro(rules)))

	rules = new(abstractfactory.USRulesAbstractFactory)

	iphone11Factory = factory.NewIPhone11Factory(rules)
	iphoneXFactory = factory.NewIPhoneXFactory(rules)

	fmt.Println("### Ordering an iPhone 11")
	newIPhone11 = iphone11Factory.OrderIPhone("standard", iphone11Factory.FuncCreateIphone())
	newIPhone11Pro = iphone11Factory.OrderIPhone("pro", iphone11Factory.FuncCreateIphone())

	fmt.Println("### Ordering an iPhone X")
	newIPhoneX = iphoneXFactory.OrderIPhone("standard", iphoneXFactory.FuncCreateIPhone())
	newIPhoneXMax = iphoneXFactory.OrderIPhone("max", iphoneXFactory.FuncCreateIPhone())

	assert.Equal(t, reflect.TypeOf(newIPhoneX), reflect.TypeOf(iphone.NewIphoneX(rules)))
	assert.Equal(t, reflect.TypeOf(newIPhoneXMax), reflect.TypeOf(iphone.NewIphoneXMax(rules)))

	assert.Equal(t, reflect.TypeOf(newIPhone11), reflect.TypeOf(iphone.NewIphone11(rules)))
	assert.Equal(t, reflect.TypeOf(newIPhone11Pro), reflect.TypeOf(iphone.NewIPhone11Pro(rules)))
}
