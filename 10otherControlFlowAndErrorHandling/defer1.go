package main

import "fmt"

/*
	the "defer" keyword is used to delay the execution
	of a part of our program for some time
*/

/*
	a little bit more detailed explanation about the "defer"
	keyword is that it delays execution of a part of our program
	till after the function it is in is done executing its last statement,
	but before the function actually ends
*/

func main() {

	//without defer
	fmt.Println("first")
	fmt.Println("second")
	fmt.Println("third")
	/*
		the order above will be:
		first,second,third
	*/

	//with defer
	fmt.Println("monday")
	defer fmt.Println("tuesday")
	fmt.Println("wednessday")
	/*
		the order above will be:
		monday,wednessday,tuesday.

		the program recognises the defered statement,
		it runs the entire program first, and then
		comes back to run to defer statements
	*/

	//important note
	/*
		defer functions execute in 'LIFO' order,"last in,first out",
		this means that if there are more than one defer statements, the last
		one will be first one to be executed
		example below
	*/
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	/*
		because of the LIFO order the outputs will be:
		three,two,one
	*/

}
