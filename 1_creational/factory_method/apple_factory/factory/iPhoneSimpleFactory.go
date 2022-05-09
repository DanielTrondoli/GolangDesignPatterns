package factory

import (
	"fmt"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/factory_method/apple_factory/model"
)

type IPhoneSimpleFactory struct{}

func (f IPhoneSimpleFactory) GetIphone11Standard() model.IIPhone {
	return model.NewIPhone11("11", "Standard")
}

func (f IPhoneSimpleFactory) GetIphone11Pro() model.IIPhone {
	return model.NewIPhone11Pro("11", "Pro")
}

func (f IPhoneSimpleFactory) GetIphoneXStandard() model.IIPhone {
	return model.NewIPhoneX("X", "Standard")
}

func (f IPhoneSimpleFactory) GetIphoneXMax() model.IIPhone {
	return model.NewIPhoneXMax("X", "Max")
}

func (f IPhoneSimpleFactory) OrderIPhone(generation, level string) (model.IIPhone, error) {
	var device model.IIPhone

	if generation == "X" {
		if level == "Standard" {
			device = model.NewIPhoneX(generation, level)
		} else if level == "Max" {
			device = model.NewIPhoneXMax(generation, level)
		}
	} else if generation == "11" {
		if level == "Standard" {
			device = model.NewIPhone11(generation, level)
		} else if level == "Pro" {
			device = model.NewIPhone11Pro(generation, level)
		}
	}

	if device == nil {
		return nil, fmt.Errorf("generation or level not exist !!!")
	}

	return device, nil
}
