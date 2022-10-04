package main

import (
	"fmt"
	"sync"
)

/*
	wait groups can also be passed and used locally
	within functions
	instead of being used globally
*/

//accept a pointer to the wait group
func sayHello(wg *sync.WaitGroup) {
	fmt.Println("hello")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	//pass in the address of created waitgroup
	go sayHello(&wg)
	wg.Wait()
}
