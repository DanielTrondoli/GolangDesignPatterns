package factory

import "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/abstract_factory/app/service"

type IServiceAbstractFactory interface {
	GetUserService() service.IUserService
	GetCarService() service.ICarService
}
