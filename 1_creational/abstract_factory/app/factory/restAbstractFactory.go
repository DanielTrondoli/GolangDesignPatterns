package factory

import "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/app/service"

type RestAbstractFactory struct {
	IServiceAbstractFactory
}

func (f RestAbstractFactory) GetUserService() service.IUserService {
	return new(service.UserRestApiService)
}

func (f RestAbstractFactory) GetCarService() service.ICarService {
	return new(service.CarRestApiService)
}
