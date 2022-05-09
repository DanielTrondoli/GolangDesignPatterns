package deepshallow_test

import (
	"fmt"
	"testing"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/prototype/deep_shallow/model"
	"github.com/stretchr/testify/assert"
)

// TestDeepShallow: Em GoLang uma atribuição ja gera um deep Copy,
// porém, quando existe um ponteiro no type nao é feita uma copia
// dos valores apontados e uma nova referencia é atribuida, é apenas
// criada uma nova variavel de ponteiro apontando para a mesma
// referencia que a variavel original.
func TestDeepShallow(t *testing.T) {

	u1 := model.User{}
	u1.Name = "Daniel"
	u1.Age = 29
	ad := model.Address{Street: "av. Jose Alfredo", Number: 30}
	u1.SetAdress(&ad)
	// clonando com atribuição normal
	u2 := u1

	assert.Equal(t, u2.Name, "Daniel")
	assert.Equal(t, u2.Age, 29)
	assert.Equal(t, u2.Address.Street, "av. Jose Alfredo")
	assert.Equal(t, u2.Address.Number, 30)

	u1.Age = 30
	u1.Address.Number = 500

	assert.Equal(t, u1.Name, "Daniel")
	assert.Equal(t, u1.Age, 30)
	assert.Equal(t, u1.Address.Street, "av. Jose Alfredo")
	assert.Equal(t, u1.Address.Number, 500)

	assert.Equal(t, u2.Name, "Daniel")
	assert.Equal(t, u2.Age, 29)
	assert.Equal(t, u2.Address.Street, "av. Jose Alfredo")
	assert.Equal(t, u2.Address.Number, 500)
	// clonando com uma função de clone
	u3 := u1.Clone()

	u1.Age = 66
	u1.Address.Number = 666

	assert.Equal(t, u1.Name, "Daniel")
	assert.Equal(t, u1.Age, 66)
	assert.Equal(t, u1.Address.Street, "av. Jose Alfredo")
	assert.Equal(t, u1.Address.Number, 666)

	assert.Equal(t, u3.Name, "Daniel")
	assert.Equal(t, u3.Age, 30)
	assert.Equal(t, u3.Address.Street, "av. Jose Alfredo")
	assert.Equal(t, u3.Address.Number, 500)

	fmt.Println(u1.Print())
	fmt.Println(u2.Print())
	fmt.Println(u3.Print())

	fmt.Println()
}
