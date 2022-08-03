package main

import "fmt"

func main() {
	/*
		in creating pointers, it is sometimes necessary to declare
		the exact type of pointer we want, so as to have more control
		over the program
	*/

	var number int = 5 //we have created a variable of type "int"
	var pNum *int = &number
	/*
		above, we have explicitly set our pointer to point to a data type of "int"
		and this pointer points to the address of "number" hence the "&number"

		Note: recall that "*int" means declare a pointer to type "int", it is not
		the same as derefrencing a value
	*/

	fmt.Println(number)
	*pNum = 16 //derefrence the pointer and change the value
	fmt.Println(number)
}
