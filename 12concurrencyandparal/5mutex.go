package main

import (
	"fmt"
	"sync"
)

/*
	to oversimplify, a mutex provides a way to
	lock certain parts of the application

	this becomes necessary when performing concurrency,
	and different goroutines are trying to access or modify
	certain parts of the application

	please refer to the below link for more info
	https://golangbot.com/mutex/#:~:text=A%20Mutex%20is%20used%20to,Mutex%20namely%20Lock%20and%20Unlock.

	consider the example below
*/

/*var value int = 0

var wg = sync.WaitGroup{}

func increment() {
	value = value + 1
	wg.Done()
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println(value)
}*/

/*
	the above program that was commented out has a race
	condition, in which one goroutine is racing againt another.
	and since one goroutine is actually modifying a resource that
	will be used in another goroutine, it leads to undesireable and
	unpredictable happenings

	here, we can use mutexes to lock a certain part of the program, so that
	only a single goroutine is working on a particular resource at any particular
	time
*/

var value int = 0

var wg = sync.WaitGroup{}

//create the mutex
var mu = sync.Mutex{}

func increment() {
	mu.Lock()
	value = value + 1
	mu.Unlock()
	wg.Done()

	/*
		the above code will basically lock the part
		of the program that is between the lock, and unlock,
		so that only one goroutine can work on that part of the
		program at a time
	*/
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println("final value is", value)
}
