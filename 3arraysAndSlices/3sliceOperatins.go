package main

import "fmt"

func main() {
	scores := []int{1, 2, 3, 4}
	//appending a value to a slice
	scores = append(scores, 5)
	/*
		the append() function takes in two arguements,
		the first one is the slice to append to, and the second
		one is the value to be appended, note, the append() function
		returns a new slice, so we should set our old slice to the value of
		this new slice with the data already appended
	*/
	fmt.Println(scores) //[1,2,3,4,5]

	//appending other slices
	a := []int{1}
	a = append(a, []int{2, 3, 4}...) //this special syntax is used to spread or open up the slice
	fmt.Println(a)                   //[1,2,3,4]

	//slice ranges
	/*
		slice ranges are simply a way to pick
		a range of values from a slice
	*/
	names := []string{"john", "doe", "jane", "foster", "bruce", "wayne"}
	rangeOne := names[1:3]
	/*
		the ":" syntax above is used to basically divide a slice
		in the above example, we are creating a new slice that
		goes from index position 1(included) to index position 3(excluded)
		of the "names" slice
	*/
	fmt.Println(rangeOne) //[doe,jane]

	rangeTwo := names[2:]
	/*
		the following above example will go from index position 2, till the end
		of the slice. this *will* however include the last element
	*/
	fmt.Println(rangeTwo) //["jane","foster","bruce","wayne"]

	rangeThree := names[:3]
	/*
		the above example will go from the start position to
		the index position 3(excluded)
	*/
	fmt.Println(rangeThree) //["john","doe","jane"]
}
