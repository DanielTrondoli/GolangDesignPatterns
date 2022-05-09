package service

import (
	"fmt"
	"strconv"
)

type UserEJBService struct {
	IUserService
}

func (u UserEJBService) Save(name string) {
	fmt.Println("Saving " + name + " through EJB's interface")
}

func (u UserEJBService) Delete(id int) {
	fmt.Println("Removing User #" + strconv.Itoa(id) + " through EJB's interface")
}
