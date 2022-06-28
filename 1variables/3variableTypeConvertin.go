package main

import "fmt"

func main() {
	var i int = 32              //declare this variable as an integer
	fmt.Printf("%v,%T\n", i, i) //print the value and type

	var f float32             //initialise another variable of type float
	f = float32(i)            //case int i into a type of float32 and assign it to f
	fmt.Printf("%v,%T", f, f) //print the value and type, notice that the type has been changed into float32

}

/*
you might wonder why it is necessary to convert from a type int, to a type
flaot before assignment, this is because in golang and similar static typed languages,
a varibale in one type cannot be set to another variable in another type.
Hench, the below conventional way will not work if subsituted for line 10
"f = i"
*/
