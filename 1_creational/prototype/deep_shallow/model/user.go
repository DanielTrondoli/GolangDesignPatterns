package model

import "fmt"

type User struct {
	Name    string
	Age     int
	Address *Address
}

func (u User) Clone() User {
	new := u
	newAddress := u.Address.Clone()
	new.Address = &newAddress
	return new
}

func (u User) GetAdress() Address {
	return *u.Address
}

func (u *User) SetAdress(a *Address) {
	u.Address = a
}

func (u User) PrintMemoryAdress() {
	fmt.Printf("\n %p", u.Address)
}

func (u User) Print() string {
	return fmt.Sprintf("%s %d %v", u.Name, u.Age, *u.Address)
}
