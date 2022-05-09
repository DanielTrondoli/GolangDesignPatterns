package housebuilder_test

import (
	"testing"

	housebuilder "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/builder/house_builder"
	"github.com/stretchr/testify/assert"
)

func TestBuildHouse(t *testing.T) {

	builder := new(housebuilder.HouseBuilder)

	house := builder.SetBathroomsNumber(4).SetBedroomsNumber(3).SetKitchens(1).SetPoolsNumber(0).Build()

	assert.Equal(t, house.BathroomsNumber, 4, "The number of Bathrooms is wrong !")
	assert.Equal(t, house.BedroomsNumber, 3, "The number of Bedrooms is wrong !")
	assert.Equal(t, house.KitchensNumber, 1, "The number of Kitchens is wrong !")
	assert.Equal(t, house.PoolsNumber, 0, "The number of Pools is wrong !")

}
