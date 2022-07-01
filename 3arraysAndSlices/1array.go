package main

import "fmt"

func main() {
	//arrays are collection types in go,
	//mostly used to store a collection of data in a single variable

	scores := [4]int{70, 32, 52, 45} //declaring an array
	/*
		one way to declare an array is to first give the array a name,
		then specify the size of the array, with the [size] syntax,
		this is the number of elements that the array can hold, then lastly, we
		specify the type of data that the array is designed to store, an array in golang
		can only store one type of data.
		in the case above, the array is an array of 4 integers.
		Then we use the {} syntax to input the values into the arrays we want
	*/
	fmt.Println(scores)
	fmt.Println(scores[3]) //we can also get at specific values from an array by using their index position

	numbers := [...]int{1, 2, 3, 4, 5, 6}
	/*
		another way to create an array without worrying about the size will be
		to use the [...] syntax, this makes will create an arrya that's just large
		enough to carry the size we give it
	*/
	fmt.Println(numbers)
	fmt.Println(numbers[2]) //getting a specific element

	var people [3]string //declare an empty array if you wish to assign elements later
	people[0] = "doe"    //assign the first element of the array to a value by specifying it's index and giving it a value
	people[1] = "john"
	fmt.Println(people)

	//getting the size of an array
	//we can get the size(length) of an array by using the built in len() function,
	//and passing in
	//the array

	cars := [...]string{"volvo", "mercedes", "audi", "tesla"}
	fmt.Println(len(cars))

	/*
		an important note about arrays in golang, when we copy an array,we
		are essentially copying the full array into another point in the memory,
		the new array is not pointing to the old array, but is instead a new array in itself
		observe below
	*/

	a := [...]int{1, 2, 3} //an array is created
	b := a                 //a variable b, is set to a, b has now become it's very own array,
	b[2] = 2               //set a value in b, notice that this will not affect any values in a
	fmt.Println(a)         //[1,2,3]
	fmt.Println(b)         //[1,2,2]
	/*
		we should therefore be careful when copying really big arrays,
		this could result in slow performance
	*/

}
