package service

import "fmt"

type CarEJBService struct {
	ICarService
}

func (c CarEJBService) Save(model string) {
	fmt.Println("Saving " + model + " car through EJB's interface")
}

func (c CarEJBService) Update(newModel string) {
	fmt.Println("Updating " + newModel + " car through EJB's interface")
}
