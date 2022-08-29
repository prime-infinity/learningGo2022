package main

import "fmt"

/*
	reveiver functions in golang is an attempt to
	mimick inheritance in traditional OOP languages.
	Receiver functions give us a way to couple or tie
	certain functions to custom objects and data types in golang.
	similar to creating methods in classes in OOP
*/

//first we create a custom type
type car struct {
	name     string
	topSpeed string
}

//second, we create the receiver function
/*
	notice that the func below accepts our custom type
	as a sort of arguement, this couples it with our
	custom type and the func cannot be accessed normally
	except through our custom type
*/
func (c car) describe() {
	fmt.Println("car is" + c.name + "and top speed is" + c.topSpeed)
}

func main() {
	aCar := car{
		name:     "Bugatti cheron",
		topSpeed: "300mph",
	}
	aCar.describe()
}
