package main

import (
	//"bytes"
	"fmt"
	//"math/rand"
	//"time"
)

func ManageStructure() {
	c := Car{}
	c.Name = "Amaze"
	c.Age = 1
	c.ModelNo = 2019
	c.LetsPrint()
	fmt.Println(c.GetCarName())
}
