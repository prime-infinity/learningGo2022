package main

import "fmt"

func main() {

	/*declaring a varibale*/

	/*in the first method below, we declare the variable
	give it a name and give it a type,in this case, an integer
	we can then assign it as needed*/
	var n int
	n = 2
	fmt.Println(n)

	/*in the second method, we can declare and assign the
	variable in the same line */
	var s int = 54
	fmt.Println(s)

	/* even simpler method, we can let golang guess the data type
	for us */
	i := 88
	fmt.Println(i)
}
