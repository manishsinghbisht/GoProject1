package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int
	c = make(chan int)
	fmt.Println("UnBuffered Channel: ", c)

	//send
	go func() {
		c <- 1
	}()

	//sniff
	val1 := <-c
	fmt.Println("Sniff first value from channel: ", val1)

	//send
	go func() {
		c <- 2
	}()
	//sniff
	time.Sleep(time.Second * 2)
	val2 := <-c
	fmt.Println("Sniff second value from channel: ", val2)
}
