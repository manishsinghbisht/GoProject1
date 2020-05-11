package main

import (
	"fmt"
	"time"
)

//This pinger waits for a ping and prints a ping on receival. Then sends a pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("Ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

//This ponger waits for a pong and prints a ping on receival. Then sends a ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("Pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)
	go pinger(ping, pong)
	go ponger(ping, pong)

	//The main go routing starts the ping/pong by sending in ping channel
	ping <- 1

	//select{} or
	for {
		//Block the main thread untill interrupt
		time.Sleep(time.Second)
	}
}
