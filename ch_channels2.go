package main

import (
	"fmt"
)

type Gun struct {
	Model string
}

func main() {
	c := make(chan int, 3)
	fmt.Println("Buffered Channel: ", c)

	//Send
	go func() {
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		c <- 6
		c <- 7
		close(c)
	}()

	//Sniff
	for i := range c {
		fmt.Println(i)
	}

	newChannel := make(chan *Gun, 3)

	//Send
	go func() {
		newChannel <- &Gun{"Benelli Black Eagle - 1991"}
		newChannel <- &Gun{"Knight MK-85 - 1985"}
		newChannel <- &Gun{"Glock 17 - 1982"}
		newChannel <- &Gun{"Ruger 10/22 - 1964"}
		newChannel <- &Gun{"Remington 1100 - 1963"}
		newChannel <- &Gun{"M-16 - 1963"}
		close(newChannel)
	}()
	//Sniff
	for i := range newChannel {
		fmt.Println(i.Model)
	}
}
