package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quit chan struct{}) {
	//Switch Case
	// Select for channel async
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("Recieved!")
		case <-quit:
			fmt.Println("Quit")
			os.Exit(0)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan struct{})
	go func() {
		c <- 1
		quit <- struct{}{}
	}()

	go Select(c, quit)
	select {}
}
