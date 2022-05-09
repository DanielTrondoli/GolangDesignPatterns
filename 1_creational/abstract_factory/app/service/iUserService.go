package service

type IUserService interface {
	Save(name string)
	Delete(id int)
}
