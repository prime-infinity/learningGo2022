package main

import "fmt"

/*
	In golang,we can embedded structs
	in one another, similar to how one would do inheritance in an OOP
	language, observe below
*/

type Car struct {
	name    string
	country string
}

type PersonCar struct {
	Car         //embedding the Car struct
	model       string
	topSpeedMph int
	carType     string
}

func main() {
	myCar := PersonCar{}
	myCar.name = "Mercedes" //as you can see, we can also access instances of the embedded classes
	myCar.country = "Germany"
	myCar.model = "SLS AMG"
	myCar.topSpeedMph = 190
	myCar.carType = "sports"

	fmt.Println(myCar) //creates and assigns all fields in all instances

	/*
		Note: the two structs are both independent structs, there is no inheritance
		here as in the case of OOP.this is a composition relationship
	*/
	/*
		Also note: in the above method, we did not have to worry ourselves
		about the internal stucture of the struct, however, if we decide to
		use the "literal syntax", we need to know how the struct was structured
		example below
	*/

	anotherCar := PersonCar{Car: Car{"BMW", "Germany"}, model: "M3", topSpeedMph: 180, carType: "Coupe"}
	fmt.Println(anotherCar)

}
