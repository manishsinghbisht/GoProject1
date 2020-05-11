package main

import (
	"fmt"
)

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("Go interfaces2...")
	Anything(2.14)
	Anything("Manish")
	Anything(struct{}{})
	mymap := make(map[string]interface{})
	mymap["name"] = "Ram"
	mymap["age"] = 40
	fmt.Println(mymap)

}
