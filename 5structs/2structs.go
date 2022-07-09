package main

import "fmt"

func main() {
	//we can also use "anonymous" stucts,
	//which is like a nameless struct we can
	//declare and use qickly with little syntax
	//observe example below

	anObject := struct {
		name string
		size int
	}{name: "Item", size: 5}
	fmt.Println(anObject)

	/*
		it has two sets of curly brackets,
		the first is paired to the "struct" keyword,
		and it defines the structure of the fields of the
		struct, the second set of brackets simply initialise and
		assign values to the fields
	*/

	/*
		Note: as oppose to maps and slices, structs are value types
		meaning that when we copy a struct, it creates its own data in memory
		for the new copy, and does not reference the old one.Hence, performing
		operations on the copy will not mutate the old one.
		example below
	*/

	anotherObject := anObject         //making a copy of the first struct
	anotherObject.name = "new object" //mutating the copy
	fmt.Println(anObject.name)        //Item
	fmt.Println(anotherObject.name)   //new object

}
