package main

import "fmt"

var n int = 27

func main() {
	fmt.Println(n) //27 will be printed here
	var n int = 76
	fmt.Println(n) //76 will be printed here
	/*this phenomenon above is known as "shadowing", in this case,
	the variable with the innermost scope is treated with priority.*/
}
