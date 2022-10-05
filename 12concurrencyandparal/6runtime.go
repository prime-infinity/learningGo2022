package main

import (
	"fmt"
	"runtime"
)

func main() {
	//The following will print out the number of
	//available threads or logical processor available for use
	fmt.Printf("threads: %v\n", runtime.GOMAXPROCS(-1))
}
