package main

import "fmt"

func main() {
	x := 0

	//simulating a while loop with for
	for x < 5 {
		//fmt.Println(x)
		x++
	}

	//traditional for loop
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	//loop through more than one variable at the same time
	for z, y := 0, 0; z < 10; z, y = z+1, y+2 {
		//fmt.Println(z, y)
	}

	//simulating a do-while loop, or a loop that runs
	//infinitely until an internal condition is met
	b := 0
	for {
		//fmt.Println(b)
		b += 2
		if b > 12 {
			break //the "break" keyword breaks out of the entire loop
		}
	}

	//a loop within a loop
	for o := 0; o <= 2; o++ {
		for i := 0; i < 2; i++ {
			//fmt.Println(o, i)
		}
	}

	//looping through a slice
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	//looping through a map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
