package model

import "fmt"

type IIPhone interface {
	Model()
	Assemble()
	Certificates()
	Pack()
	GetHardware()
}

type AbstracteIPhone struct {
	generation string
	level      string
}

func (i *AbstracteIPhone) InputModel(generation, level string) {
	i.generation = generation
	i.level = level
}

func (i AbstracteIPhone) Assemble() {
	fmt.Println("Assembling all the hardwares")
}

func (i AbstracteIPhone) Certificates() {
	fmt.Println("Testing all the certificates")
}

func (i AbstracteIPhone) Pack() {
	fmt.Println("Packing the device")
}

func (i AbstracteIPhone) Model() {
	fmt.Println("IPhone " + i.generation + " - " + i.level)
}

//func (i AbstracteIPhone) GetHardware() {}
