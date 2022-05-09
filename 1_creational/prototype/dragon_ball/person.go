package dragonball

import "fmt"

type Person struct {
	Name    string
	IsAlive bool
}

func (p Person) Resurect() (Person, error) {
	if p.IsAlive {
		return Person{}, fmt.Errorf("This person is alive yet !")
	}

	return Person{
		Name:    p.Name + "_ressurected",
		IsAlive: true,
	}, nil
}
