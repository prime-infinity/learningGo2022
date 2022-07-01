package main

import "fmt"

func main() {
	//slices are similar to arrays, but are more flexible
	//declaring a slice below
	scores := []int{2, 3, 4, 5}
	/*
		a slice is declared similar to an array, but here,
		we do not need to predefine the size
	*/
	scores[2] = 3 //we can also change any value of the slice
	fmt.Println(scores)

	/*
		an important note, slices are naturally reference types, what does that mean?,
		well, it means that if you copy a slice from another slice, the copied slice does
		not create its own data in memory, but points(reference) to the same data that as the original
		slice, thus, changing or modifying a copied slice, will modify the original slice

		example below
	*/

	a := []int{1, 2, 3, 4}
	b := a         //b is referenced to a
	b[0] = 0       //modifying b will modify a
	fmt.Println(a) //[0,2,3,4] //data modified
	fmt.Println(b) //[0,2,3,4]

}
