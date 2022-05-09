package service

import "fmt"

type CarRestApiService struct {
	ICarService
}

func (c CarRestApiService) Save(model string) {
	fmt.Println("Saving " + model + " car through Rest's interface")
}

func (c CarRestApiService) Update(newModel string) {
	fmt.Println("Updating " + newModel + " car through Rest's interface")
}
