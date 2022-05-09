package model

import "fmt"

type Pasta struct {
	//Mandatory
	size     Size
	toppings []string
	sauces   []string
	cheese   bool
	pepper   bool
}

func NewPasta(size Size, cheese, pepper bool, toppings, souces []string) (p Pasta) {
	p.cheese = cheese
	p.pepper = pepper
	p.sauces = souces
	p.size = size
	p.toppings = toppings
	return p
}

func (p Pasta) PrintPasta() {
	fmt.Printf("Pasta Order: \n")
	fmt.Printf("	Size: %s\n", p.size)
	fmt.Printf("	With Cheese: %s\n", boolToYesOrNo(p.cheese))
	fmt.Printf("	With Pepper: %s\n", boolToYesOrNo(p.pepper))
	fmt.Printf("	Souces: %v\n", p.sauces)
	fmt.Printf("	Toppings: %v\n", p.toppings)
}

func boolToYesOrNo(test bool) string {
	if test {
		return "Yes"
	}
	return "No"
}
