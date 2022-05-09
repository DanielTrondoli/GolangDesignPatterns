package personanimation

import "fmt"

type Person struct {
	Name string
	L1   string
	L2   string
	L3   string
	L4   string
}

func (p Person) Clone() Person {
	newPerson := p
	newPerson.Name = p.Name + "_clone"
	return newPerson
}

func (p *Person) MoveLerft() {
	p.L1 = p.L1[1:len(p.L1)]
	p.L2 = p.L2[1:len(p.L2)]
	p.L3 = p.L3[1:len(p.L3)]
	p.L4 = p.L4[1:len(p.L4)]
}

func (p *Person) MoveRight() {
	p.L1 = " " + p.L1
	p.L2 = " " + p.L2
	p.L3 = " " + p.L3
	p.L4 = "_" + p.L4
}

func (p Person) PrintPerson() {
	fmt.Println(p.L1)
	fmt.Println(p.L2)
	fmt.Println(p.L3)
	fmt.Println(p.L4)
}
