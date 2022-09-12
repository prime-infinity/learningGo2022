package main

import "fmt"

//INTERFACES NEEDS MORE RESEARCH

/*
	interfaces are a way to group stuff based on behaviour, or methods
	similar to how structs group based on data
	interfaces are types, such as structs
*/

//first, lets create some types we want to group together
type boy struct {
	heightFeet float64
	weightKg   float64
}

type girl struct {
	heightFeet float64
	weightKg   float64
}

//secondly,we create receiver functions for both
//boy and girl structs

//boy methods
func (b boy) bmi() float64 {
	heightS := b.heightFeet * b.heightFeet
	return b.weightKg / heightS
}

func (b boy) fatPerc() float64 {
	//an hypothetical body fat percentage
	return b.heightFeet * b.weightKg / 100
}

//notice that both the girl and boy have the same methods

//girl methods
func (g girl) bmi() float64 {
	heightS := g.heightFeet * g.heightFeet
	return g.weightKg / heightS
}

func (g girl) fatPerc() float64 {
	//an hypothetical body fat percentage
	return g.heightFeet * g.weightKg / 100
}

/*
	instead of creating new methods every time for
	both boy and girl, we can group them together using
	their methods
*/

//we create an interfac

type gender interface {
	bmi() float64
	fatPerc() float64
}

/*
	the above creates an interface.
	in order for data to be in the interface, they *must* have two
	methods "bmi" and "fatPerc" associated with them.
	notice that the girl and boy structs have both methods associated with them
*/

//we can now create a generic function that works for both types

//the below function takes in the "gender" type, and
//we can access the respective methods on the types inside the
//gender interface type
func printGenderInfo(g gender) {
	fmt.Printf("BMI of %T is: %0.2f \n", g, g.bmi())
	fmt.Printf("fat percentage of %T is : %0.2f \n", g, g.fatPerc())
}

func main() {

	//the below part is not
	//too clear yet, further
	//revision is needed

	genders := []gender{
		boy{heightFeet: 6, weightKg: 76},
		girl{heightFeet: 5, weightKg: 55},
	}

	for _, v := range genders {
		printGenderInfo(v)
		fmt.Println("---")
	}
}
