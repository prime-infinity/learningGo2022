package main

import "fmt"

func main() {
	studentScores := map[string]int{
		"john": 78,
		"doe":  54,
		"jane": 89,
		"anna": 53,
	}
	fmt.Println(studentScores)

	//we can view a particular value from the map
	//with the below syntax, similar to getting values
	//from an array, we put the "key" of the value to get in square brackets
	//[key]
	fmt.Println(studentScores["john"]) // 78

	//we can also add some data into the map by setting a key and assigning it
	//to something example below
	studentScores["newstudent"] = 85 //this key - value pair has been added to the map

	/*
		Note: the return order of a map is not fixed, that is, the elements
		of the map loose their order when used in the program.so even if the
		first key - value pair is "john":78, it may not be stored that way in memory
	*/

	//to delete an element from a map, we use the built in "delete()" function,
	//we pass in *two* arguements, the first is the map we want to delete from,
	//and the second is the "key" of the value to delete e.g
	delete(studentScores, "anna") //deletes "anna" from the map

	//to check if a particular "key" exist in our map,we use the below syntax
	_, isPresent := studentScores["oldStudent"]
	fmt.Println(isPresent) //false

	//as in arrays and slices, we can also view the length of a map with
	//the built in "len()" function
	fmt.Println(len(studentScores)) //4

	/*
		another very important note is that, like slices, maps in golang are passed by reference
		which means that when you copy a map, it does not create a new map, but instead the new map
		is pointing(referencing) to the old map, so manipulating any values of the new map will also
		mutate the old map. e.g
	*/

	newSlice := studentScores
	delete(newSlice, "doe")
	fmt.Println(newSlice)
	fmt.Println(studentScores) //notice, "doe" is also removed from studentScores too
}
