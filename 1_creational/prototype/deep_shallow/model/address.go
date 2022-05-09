package model

type Address struct {
	Street string
	Number int
}

func (a Address) Clone() Address {
	new := new(Address)
	new.Number = a.Number
	new.Street = a.Street
	return *new
}
