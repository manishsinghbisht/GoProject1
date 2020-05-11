package main

import (
	//"bytes"
	"fmt"
	//"math/rand"
	//"time"
)

func (x Car) LetsPrint() {
	fmt.Println(x)
}

func (x Car) GetCarName() string {
	return x.Name
}
