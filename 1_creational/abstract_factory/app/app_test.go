package app_test

import (
	"fmt"
	"testing"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/app/factory"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/app/service"
)

func TestAppAbstractFactory(t *testing.T) {

	executeServices(new(factory.EJBAbstractFactory))
	executeServices(new(factory.RestAbstractFactory))

}

func executeServices(factory factory.IServiceAbstractFactory) {
	var userService service.IUserService
	var carService service.ICarService

	fmt.Println("#-------------------------------")
	carService = factory.GetCarService()
	carService.Save("Camaro Amarelo")
	carService.Update("BMW X5")

	fmt.Println("#-------------------------------")
	userService = factory.GetUserService()
	userService.Save("Daniel Trondoli")
	userService.Delete(1)

}
