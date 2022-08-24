package main

import "fmt"

/*
	functions can also return data from them.
	to make them return data, we simply specify
	the type of data we are expecting the function to
	return after before the "{" bracket
	example below
*/

//this func is expected to return a float64 value
func double(num float64) float64 {
	return num * 2
}

func squareArea(l float64, b float64) float64 {
	//remember how to accept multiple args
	return l * b
}

//returning multiple data values
/*
	in returning multiple data types,
	we simply put the expected return data types in
	brackets and seperate them with a "," as shown below
*/
//the below function accepts a slice of strings
//and return the first and last items
func giveFirstAndLastOfStringSlice(s []string) (string, string) {
	first := s[:1]
	last := s[len(s)-1:]

	//we then return then with a "," seperated
	return first[0], last[0]
}

//update..24 august
/*
	we can also return pointers to values
	in the below function, we are returning a
	pointer to an integer
*/

func returnSome(value int) *int {
	return &value
	//return the address of the value
}

func main() {
	number := double(15.3)
	fmt.Println(number)

	area := squareArea(7, 3)
	fmt.Println(area)

	theSlice := []string{"one", "two", "three", "four", "five", "six"}
	//accepting multiple return values with a "," seperated assignment
	first, last := giveFirstAndLastOfStringSlice(theSlice)
	fmt.Println(first, last)

	//update..24 august
	value := returnSome(5)
	fmt.Println(*value) //derefrenced

}
