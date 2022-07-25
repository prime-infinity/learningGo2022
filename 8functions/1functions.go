package main

import "fmt"

/*
	we use functions to divide the overall program
	into smaller,reusable and
	easier to understand parts
*/

//declaring a simple function
/*
	we declare a function by using the "func" keyword,
	followed by the name of the function, then "{}" brackets
	that will hold the code to run when the function is called
*/

func sayHi() {
	//code to run
	fmt.Println("i am alive, feed me your bytes")
}

func main() {
	//call the function
	sayHi() //"i am alive, feed me your bytes"
}
