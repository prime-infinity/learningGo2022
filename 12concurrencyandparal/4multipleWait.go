package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func sayName(n string) {
	fmt.Println(n)
	wg.Done()
}

func main() {
	names := []string{"osamede", "richmond", "osasu"}
	wg.Add(len(names))
	for _, v := range names {
		go sayName(v)
	}
	wg.Wait()
}
