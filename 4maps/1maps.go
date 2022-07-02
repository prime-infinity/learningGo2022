package main

import "fmt"

func main() {
	/*
		a map is similar to objects in javascript,
		basically a collection data type having key-value pairs,
		with keys to the left and values to the right.
		an example below
	*/

	studentScores := map[string]int{
		"john": 78,
		"doe":  54,
		"jane": 89,
		"anna": 53,
	}

	/*
		maps in go are a little bit different though,
		we must always specify the type of the keys and the type
		of the values that are to be assigned to the keys, in the example above,
		we are creating a map of keys of type "string" and mapping them to values of type "int",
		hence the syntax to initialise the map, map[string]int
	*/
	fmt.Println(studentScores)

	//if the key - value pairs are not present at the time of
	//creating the map, we can also use the "make" syntax to initialise an empty map

	humanProps := make(map[string]string)
	humanProps = map[string]string{
		"hair": "black",
		"eyes": "black",
		"skin": ".....black",
	}

	fmt.Println(humanProps)
}
