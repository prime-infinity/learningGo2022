package main

import "fmt"

/*
	the oversimplified explanation about pointers is that:
	they are a data type on their own, all they do is point to
	another location in memory where another value is stored.
*/

func main() {

	//first of all, lets declare a variable
	number := 6

	/*
		the above value is stored in memory, to get the
		address of its location in memory,
		we put a "&" before the value.
	*/

	fmt.Println(&number)
	/*
	 the above will print out a weird stuff like "0xc000012044",
	 do not fret, this is actually the location of the value in your
	 computer memory
	*/

	pNum := &number
	/*
		above, we are setting another variable to be equal to the
		address of "number"
	*/

	/*
		if we have the address of a value, to get back the value, we need
		"dereference" the address by putting a "*" before the value
	*/

	fmt.Println(*pNum) //6

	/*
		we can also directly perform mutations on the pointers of values
		by first derefrencing them and mutating them
		and this will translate to peforming mutations on the values themselves
		example below
	*/

	*pNum = 13
	fmt.Println(number) //13

	//do more research on the importance of using pointers
}
