package pastabuilder_test

import (
	"testing"

	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/builder/pasta_builder/builder"
	"github.com/DanielTrondoli/GolangDesignPatterns/1_creational/builder/pasta_builder/model"
	"github.com/stretchr/testify/assert"
)

func TestPastaBuilder(t *testing.T) {
	pasta := builder.NewPastaBuilder(model.STANDARD).AddSouce("Tomato").WithCheese().AddToppings("Bacon").AddSouce("Chedar").Order()
	pasta.PrintPasta()

	pasta2 := model.NewPasta(model.STANDARD, true, false, []string{"Bacon"}, []string{"Tomato", "Chedar"})
	pasta2.PrintPasta()
	assert.Equal(t, pasta2, pasta, "Objects are not equivalent")
}
