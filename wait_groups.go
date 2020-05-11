package main

import (
	"fmt"
	"sync"
)

func main() {
	//wait groups for add, done, wait
	wg := &sync.WaitGroup{}

	//Use Mutexes - for accessing a shared resource (critical)

	wg.Add(2)
	//Anonymous go func
	go func() {
		fmt.Println("Func 1")
		wg.Done()
	}()

	//Anonymous go func
	go func() {
		fmt.Println("Func 2")
		wg.Done()
	}()

	fmt.Println("Finish")
	wg.Wait()
	fmt.Println("Exit")
}
