package builder

import "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/builder/pasta_builder/model"

type pastaBuilder struct {
	size     model.Size
	toppings []string
	sauces   []string
	cheese   bool
	pepper   bool
}

func NewPastaBuilder(size model.Size) pastaBuilder {
	builder := new(pastaBuilder)
	builder.size = size
	builder.cheese = false
	builder.pepper = false
	builder.sauces = make([]string, 0)
	builder.toppings = make([]string, 0)
	return *builder
}

func (p pastaBuilder) AddSouce(souce string) pastaBuilder {
	p.sauces = append(p.sauces, souce)
	return p
}

func (p pastaBuilder) AddToppings(toppings string) pastaBuilder {
	p.toppings = append(p.toppings, toppings)
	return p
}

func (p pastaBuilder) WithCheese() pastaBuilder {
	p.cheese = true
	return p
}

func (p pastaBuilder) WithPepper() pastaBuilder {
	p.pepper = true
	return p
}

func (p pastaBuilder) Order() model.Pasta {
	return model.NewPasta(p.size, p.cheese, p.pepper, p.toppings, p.sauces)
}
