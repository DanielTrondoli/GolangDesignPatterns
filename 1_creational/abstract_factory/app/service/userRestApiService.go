package service

import (
	"fmt"
	"strconv"
)

type UserRestApiService struct {
	IUserService
}

func (u UserRestApiService) Save(name string) {
	fmt.Println("Saving " + name + " through Rest's interface")
}

func (u UserRestApiService) Delete(id int) {
	fmt.Println("Removing User #" + strconv.Itoa(id) + " through Rest's interface")
}
