package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

/*
	it is generally good practice to have channels
	dedicated to either sending or receiving

	the way we do this is below
*/

func main() {
	//syntax to make a channel
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) {
		//we make this channel to be send only
		//by putting the args in the brackets

		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		//this channel is receive only, note the
		//position of the arrow

		ch <- 42
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
