package main

import (
	//"bytes"
	"fmt"
	"strings"
	//"math/rand"
	//"time"
)

func DoString() {
	fmt.Println("Strings.....")
	var m1 string
	m1 = "My first string!"
	fmt.Println(m1)

	m2 := "My second string!"
	fmt.Println(m2)

	fmt.Println("Do string1 contains string2: ", strings.Contains(m1, m2))
	fmt.Println("Replace M in string1 with Y: ", strings.ReplaceAll(m1, "M", "Y"))
	fmt.Println("Split string in string2, retruns array: ", strings.Split(m2, " "))

}
