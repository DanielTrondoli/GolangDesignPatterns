package service

type ICarService interface {
	Save(model string)
	Update(newModel string)
}
