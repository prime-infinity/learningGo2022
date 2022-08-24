package main

/*
	a panic occurs in a go when the program
	can no longer continue and enters a state of "panic".
	a panic may be thought of as execptions in other languages
*/

func main() {

	//create a panic
	//we can create our own panic with custom message like below
	panic("someting bad has happened")
	//note that calling a panic will immediately stop the
	//current goroutine or function it is called in

	/*a, b := 1, 0
	ans := a / b
	fmt.Println(ans)*/
}
