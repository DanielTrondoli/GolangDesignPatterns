package dragonball_test

import (
	"testing"

	dragonball "github.com/DanielTrondoli/GolangDesignPatterns/1_creational/prototype/dragon_ball"
)

func TestResurectPerson(t *testing.T) {

	krillin := dragonball.Person{Name: "Krillin", IsAlive: true}

	kill(&krillin)

	shenlong := dragonball.Shenlong{}

	krillin_res, err := shenlong.Resurect(krillin)
	if err != nil {
		t.Fatal(err)
	}

	if !krillin_res.IsAlive {
		t.Fatal("Test 2: error, person still dead")
	}

	if krillin_res.Name != krillin.Name+"_ressurected" {
		t.Fatal("Test 1: error name not have ressurected")
	}

	krillin_res_again, err := shenlong.Resurect(krillin_res)
	if err != nil {
		t.Log(err)
	}
	if (krillin_res_again != dragonball.Person{}) {
		t.Fatal("Test 3: you can't ressurect one person alive")
	}
}

func kill(p *dragonball.Person) {
	p.IsAlive = false
}
