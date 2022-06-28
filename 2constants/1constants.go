package main

import "fmt"

const c int16 = 45

func main() {

	const aConst int = 53
	//a contant is created similar to a variable, but a "const" keyword is used instead of "var"
	//a constant *must* be available at compile time, it cannoto be set at run time
	//aConst = 4 //this will not work, constants cannot be reassigned but they *can* shadowed

	const c float32 = 45.2 // this constant is shadowed, with this inner declaration taking priority
	fmt.Println(c)

	//constants can be made up of any primitives
	const i int = 14
	const s string = "hello"
	const b bool = true
	fmt.Println(i, s, b)

	const h = 65 //this syntax can also be used where we let the compiler guess the type
}
