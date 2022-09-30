package main

import "fmt"

/*
	simply, a goroutine is the go version of an execution thread.
	it is like a virtual thread that go manages for us.
	goroutines allows us to perform concurrency.
	goroutines can execute concurrently with one another
*/

//create a simple function
func greet() {
	fmt.Println("hello")
}

func main() {
	//call the function normally
	//greet()

	//to call the function concurrently,
	//simply add the "go" keyword before calling the function

	go greet()

	/*
		note: when the code above is run, it will not print anything out.
		this does not mean it did not call the function, it did, but because of
		the nature of concurrency, the main function exits as soon as the concurrent
		function is started and does not wait for the concurrent function to finish

		we shall deal with this in the next lesson
	*/
}
