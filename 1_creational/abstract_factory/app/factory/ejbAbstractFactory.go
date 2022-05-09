package factory

import "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/app/service"

type EJBAbstractFactory struct {
	IServiceAbstractFactory
}

func (f EJBAbstractFactory) GetUserService() service.IUserService {
	return new(service.UserEJBService)
}

func (f EJBAbstractFactory) GetCarService() service.ICarService {
	return new(service.CarEJBService)
}
