package main

/*
	every go file is part of a package
 	"package main" tells the compiler that this particular
 	package should compile as an executable program.
*/

import "fmt"

/*
	here, a package is being imported from the go standard
	library to be used in our program.
	the standard library is included when you install go and
	contains packages that assist and extend functionality.
	the "fmt" package is for formatting strings and outputing
	to the console.
*/

func main() {

	/*
		the "main()" function is the entry point of the program and is
		fired automatically.
		It should be called "main()" and nothing else, there should
		only be one "main()" function in the entire program
	*/

	fmt.Println("I am alive, feed me your bytes")

	/*
		"Println()" is a method from the "fmt" package.
		it simply prints to the console
	*/
}
