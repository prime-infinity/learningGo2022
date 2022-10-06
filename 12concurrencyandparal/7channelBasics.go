package main

import (
	"fmt"
	"sync"
)

/*
	to oversimplify, channels are a way for goroutines
	to communicate with one another
*/

var wg = sync.WaitGroup{}

func main() {
	//syntax to make a channel
	ch := make(chan int)
	/*
		the above syntax is used to make a channel, set the channel
		to a variable and use the "make" keyword, always remember to
		specify the type of data that is expected to flow through a channel
		as a second arguement when creating channels,
	*/

	//create two goroutines below
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}
