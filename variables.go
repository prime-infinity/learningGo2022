package main

import "fmt"

/* variables can also be declared at the package level, here
the ":=" method cannot be used */

var a float32 = 24

/*another thing that can be done at package level
is a variable block as below*/
var (
	studentName  string = "john stark"
	studentGrade int    = 4
)

var Q string = "this is exported" //this variable can be accessed outside this package

func main() {

	/*declaring a varibale*/

	/*in the first method below, we declare the variable
	give it a name and give it a type,in this case, an integer
	we can then assign it as needed*/
	var n int
	var nameOne string
	nameOne = "john" //remember, strings in golang are doublequotes, not singlequotes
	n = 2
	fmt.Println(n, nameOne)

	/*in the second method, we can declare and assign the
	variable in the same line */
	var s int = 54
	var nameTwo = "jane"
	fmt.Println(s, nameTwo)

	/* even simpler method, we can let golang guess the data type
	for us, this below method cannot be used outside of a function*/
	i := 88
	nameThree := "jose"
	fmt.Println(i, nameThree)

	////variable delclared at the package level
	fmt.Println(a)
	//variables from variable block above
	fmt.Println(studentName, studentGrade)

	fmt.Println(Q)
}

/*
remember!, when variables are declared
at the package level and are named with uppercase as in line 17,
they are exported and can be used in other packages
*/
