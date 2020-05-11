package main

import (
	//"bytes"
	"fmt"
	//"math/rand"
	//"time"
)

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func main1() {
	var i int
	i = 15
	ptr := &i
	fmt.Println("Hex Address: ", ptr)
	fmt.Println("Value At: ", *ptr)
	m1, m2 := 2, 3
	fmt.Println("Before Swap", m1, m2)
	swap(&m1, &m2)
	fmt.Println("After Swap", m1, m2)
}
