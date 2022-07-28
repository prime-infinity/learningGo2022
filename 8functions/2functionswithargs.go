package main

import "fmt"

/*
we can also pass in arguements into functions,
in passing arguements, we give it any name we want it
to be called inside the function and give it a specific type,
example below
*/

//here, the function takes in an arguement
//called "word" of type string, this means that this
//arguement will only be accepted if it is a string
func say(word string) {
	fmt.Println(word) //prints out the "word" arguement passed to it
}

//using multiple arguements
func sayMultiple(word string, number int) {
	//we can accept multiple arguements
	//by seperating them like above
	fmt.Println(word, number)
}

//our functions can also take in other functions
//as arguements
func cycleWords(words []string, fon func(string)) {
	for _, v := range words {
		//the below code might seem strange
		//but all it does is call the function "fon"
		//for each element of the slice
		fon(v)
	}

}

func main() {

	//we call the function with its string arguement
	say("hey")
	say("your")
	//we can then recall the function multiple times with different arguements
	say("this")

	//calling a function with multiple arguements
	sayMultiple("number", 5)

	words := []string{"my", "name", "is", "omni"}

	cycleWords(words, say)
}
