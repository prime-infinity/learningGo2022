package main

import "fmt"

func main() {
	/*
		a switch statement can be thought of
		as a less verbose and more elegant way to
		combine if-else statements
	*/

	value := 3

	switch value {
	case 1:
		fmt.Println("the value is 1")
	case 2:
		fmt.Println("the value is 2")
	case 3:
		fmt.Println("the value is 3")
	default:
		fmt.Println("not 1 or 2 or 3")
	}

	/*
		the syntax for the switch statement is shown above,
		the "value" proceeding the "switch" is any value that we want the program to compare against.
		then we have a series of cases and the value proceeding the cases is what
		the program is going to compare against our "value".For that case that matches the value,
		the statements or code after the particular case will be executed.
		the "default" case will be executed if no other case is executed
	*/

	//testing multiple cases
	/*
		we can also use one switch case to check for more
		that one condition, details below
	*/
	value = 4
	switch value {
	case 1, 2, 3:
		println("one two, or three")
	case 4, 5, 6:
		println("four five or six")
	}

	//tagless method of writing switch statements
	n := 5
	switch {
	case n < 1:
		fmt.Println("n is less that 5")
	case n > 2 && n < 4:
		fmt.Println("n is greater than 2 but less thatn 4")

	}

}
