package main

import "fmt"

/*
	no need to overcomplicate things, a struct is simply
	a data type that we can use to hold related data of *any* type.
	think of a struct as a really flexible container.
*/

/*
	a struct can also be used as a blueprint, or a structure so to speak,
	by which to store a particular entity and its properties
*/

type Person struct {
	name    string
	age     int
	friends []string
	isFired bool
	scores  map[string]string
}

/*
	a struct is declared above, we simply give it a name,
	in this case, "Person",tell golang that it is a custom "type",
	then we declare the fields and their types
	we want the struct to hold.
*/

func main() {
	/*
		to use the struct, we simply assign our custom struct,
		in this case "Person" to a variable, then we assign the corresponding
		field names with values, as shown below
	*/

	aPerson := Person{
		name:    "prime",
		age:     100,
		friends: []string{"blackholes", "pulsars", "neutron stars", "the big bang"},
		isFired: true,
		scores: map[string]string{
			"brightness": "5000 suns",
			"mass":       "5000000000000 suns",
			"gravity":    "infinite",
		},
	}

	fmt.Println(aPerson)
	//we can also ask for different fields with the dot "." syntax
	fmt.Println(aPerson.scores)
	fmt.Println(aPerson.friends[len(aPerson.friends)-1]) //the big bang
}
