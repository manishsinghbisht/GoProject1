package main

import (
	//"bytes"
	"fmt"
	//"math/rand"
	//"time"
)

func DoArray() {
	fmt.Println("Arrays.....")
	arr1 := []int{1, 2, 3, 4}
	arr2 := []string{"Hi", "my", "name"}

	arr2 = append(arr2, "is Manish!")
	fmt.Println(arr1, arr2)
}
