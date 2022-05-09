package housebuilder

type HouseBuilder struct {
	bathroomsNumber   int
	bedroomsNumber    int
	livingRoomsNumber int
	kitchensNumber    int
	poolsNumber       int
}

func (d HouseBuilder) SetBathroomsNumber(n int) HouseBuilder {
	d.bathroomsNumber = n
	return d
}

func (d HouseBuilder) SetBedroomsNumber(n int) HouseBuilder {
	d.bedroomsNumber = n
	return d
}

func (d HouseBuilder) SetRoomsNumber(n int) HouseBuilder {
	d.livingRoomsNumber = n
	return d
}

func (d HouseBuilder) SetKitchens(n int) HouseBuilder {
	d.kitchensNumber = n
	return d
}

func (d HouseBuilder) SetPoolsNumber(n int) HouseBuilder {
	d.poolsNumber = n
	return d
}

func (d HouseBuilder) Build() House {
	return House{
		BathroomsNumber:   d.bathroomsNumber,
		BedroomsNumber:    d.bedroomsNumber,
		LivingRoomsNumber: d.livingRoomsNumber,
		KitchensNumber:    d.kitchensNumber,
		PoolsNumber:       d.poolsNumber,
	}
}
