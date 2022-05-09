package personanimation_test

import (
	"fmt"
	"testing"
	"time"

	personanimation "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/prototype/person_animation"
)

func TestPersonAnimation(t *testing.T) {

	sample := new(personanimation.PersonSample)
	sample.LoadFatMan()

	animation := createAnimation(sample.Sample)
	animate(animation)

	s2 := sample

	fmt.Printf("\n%p\n", &sample)
	fmt.Printf("%p\n", &s2)

	fmt.Printf("\n%p\n", &sample.Sample)
	fmt.Printf("%p\n", &s2.Sample)

}

func animate(p []personanimation.Person) {
	fmt.Println("*********************************")
	fmt.Println("*                               *")
	fmt.Println("* Adjust your screen's height!  *")
	fmt.Println("*                               *")
	fmt.Println("*********************************")
	time.Sleep(3000 * time.Millisecond)
	for _, v := range p {
		fmt.Println(v.Name)
		v.PrintPerson()
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("*********************************")
	fmt.Println("*                               *")
	fmt.Println("*             The End!          *")
	fmt.Println("*                               *")
	fmt.Println("*********************************")

}

func createAnimation(p personanimation.Person) []personanimation.Person {

	anim := make([]personanimation.Person, 0)
	anim = append(anim, p.Clone())

	aux := p.Clone()
	aux.MoveLerft()
	anim = append(anim, aux.Clone())
	aux.MoveLerft()
	anim = append(anim, aux.Clone())
	aux.MoveLerft()
	anim = append(anim, aux.Clone())
	aux.MoveRight()
	anim = append(anim, aux.Clone())
	aux.MoveRight()
	anim = append(anim, aux.Clone())
	aux.MoveRight()
	anim = append(anim, aux.Clone())
	aux.MoveRight()
	anim = append(anim, aux.Clone())
	aux.MoveRight()
	anim = append(anim, aux.Clone())
	aux.MoveRight()
	anim = append(anim, aux.Clone())
	aux.MoveRight()
	anim = append(anim, aux.Clone())
	aux.MoveLerft()
	anim = append(anim, aux.Clone())
	aux.MoveLerft()
	anim = append(anim, aux.Clone())
	aux.MoveLerft()
	anim = append(anim, aux.Clone())

	return anim
}
