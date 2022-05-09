package applefactory_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/factory_method/apple_factory/factory"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/factory_method/apple_factory/model"
	"github.com/stretchr/testify/assert"
)

func TestAppleFactory(t *testing.T) {

	factory := new(factory.IPhoneSimpleFactory)
	var newIPhone model.IIPhone
	var err error

	newIPhone, err = factory.OrderIPhone("11", "Standard")
	if err != nil {
		t.Error(err.Error())
	} else {
		printIphone(newIPhone)
	}

	newIPhone, err = factory.OrderIPhone("11", "Pro")
	if err != nil {
		t.Error(err.Error())
	} else {
		printIphone(newIPhone)
	}

	newIPhone, err = factory.OrderIPhone("X", "Standard")
	if err != nil {
		t.Error(err.Error())
	} else {
		printIphone(newIPhone)
	}

	newIPhone, err = factory.OrderIPhone("X", "Max")
	if err != nil {
		t.Error(err.Error())
	} else {
		printIphone(newIPhone)
	}

	_, err = factory.OrderIPhone("12", "Max")
	if assert.Error(t, err) {
		assert.Equal(t, "generation or level not exist !!!", err.Error())
	} else {
		printIphone(newIPhone)
	}

	newIPhone, err = factory.OrderIPhone("11", "Standard")
	if err != nil {
		t.Error(err.Error())
	} else {
		printIphone(newIPhone)
	}

	newIPhone, err = factory.OrderIPhone("11", "Pro")
	if err != nil {
		t.Error(err.Error())
	} else {
		printIphone(newIPhone)
	}

	newIPhone = factory.GetIphoneXStandard()
	assert.Equal(t, reflect.TypeOf(newIPhone), reflect.TypeOf(new(model.IPhoneX)))
	newIPhone = factory.GetIphoneXMax()
	assert.Equal(t, reflect.TypeOf(newIPhone), reflect.TypeOf(new(model.IPhoneXMax)))
	newIPhone = factory.GetIphone11Standard()
	assert.Equal(t, reflect.TypeOf(newIPhone), reflect.TypeOf(new(model.IPhone11)))
	newIPhone = factory.GetIphone11Pro()
	assert.Equal(t, reflect.TypeOf(newIPhone), reflect.TypeOf(new(model.IPhone11Pro)))
}

func printIphone(iphone model.IIPhone) {
	fmt.Println("#----------------------------")
	iphone.Model()
	iphone.Assemble()
	iphone.Certificates()
	iphone.Pack()
	iphone.GetHardware()
}
