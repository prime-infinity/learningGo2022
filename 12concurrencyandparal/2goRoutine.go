package main

import (
	"fmt"
	"sync"
)

/*
	we can use a wait group to synchronize goroutines.
	a wait group waits for goroutines to finish
*/
//create a wait group
var wg = sync.WaitGroup{}

func greet() {
	fmt.Println("hello")
	wg.Done()
}

func main() {

	/*
		wait group usage is kinda difficult to explain with code
		it is advised to do more research on your own, but i shall try
		to explain it for you guys.
	*/

	/*
		basically, ....we need to manually add a goroutine to the wait group
		with the "Add()" method on the waitgroup, then inside the function that we
		want to be concurrent, we put the "Done()" method after all the codes, this
		method tells the wg that the function is done executing and it decrements
		the wg back to zero "0", we then use the "Wait()" keyword after the concurrent
		code to make the whole program wait till its done
	*/

	wg.Add(1)
	go greet()
	wg.Wait()
}
