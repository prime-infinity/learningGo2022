package main

import "fmt"

func main() {
	a := 3 + 2i
	b := 4 + 5.2i

	fmt.Println(a + b)
	fmt.Println(a / b)
	fmt.Println(a * b)

	c := imag(a)
	fmt.Println(c)
}
