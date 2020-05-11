package main

import (
	"fmt"
)

// ICar Everything here
// will be considered
// a block comment
type ICar interface {
	Drive()
	Stop()
}

// Maruti Everything here
// will be considered
// a block comment
type Maruti struct {
	MarutiModel string
}

// Chevy Everything here
// will be considered
// a block comment
type Chevy struct {
	ChevyModel string
}

func (m *Maruti) Stop() {
	fmt.Println("Maruti stopped")
}

func (m *Maruti) Drive() {
	fmt.Println("Maruti start")
}

func (m *Chevy) Stop() {
	fmt.Println("Chevy stopped")
}

func (m *Chevy) Drive() {
	fmt.Println("Chevy start")
}

func MyCustomCar(arg string) ICar {
	return &Maruti{arg}
}

func main() {
	m := Maruti{"Maruti Alto"}
	c := Chevy{"Chevrolet Beat"}
	customModel := MyCustomCar("MyCar")

	m.Drive()
	m.Stop()
	c.Drive()
	c.Stop()
	customModel.Drive()
	customModel.Stop()
}
