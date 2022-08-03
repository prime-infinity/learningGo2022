package main

import "fmt"

/*
	passing pointers into functions
*/

/*
	the below function accepts a pointer to
	an integer, it derefrences it and performs mutations
	on it
*/
/*
	in the below function also, whenever we see "*" followed by
	any type, in the arguements of a function, eg "*int","*float64","*string",
	we are simply telling the function to accept a pointers to that type.

	thus, the below function is accepting a pointer to a type "int"
*/
func changeNumber(n *int) {
	//once inside the function, we must also derefrence the pointer also
	*n = 564 //mutates our original value,
}

func main() {

	number := 5

	pNum := &number //creating a pointer to "number"

	fmt.Println(number) //5

	changeNumber(pNum)

	fmt.Println(number) //564

	/*
		observe that the value has been mutated
	*/

}
