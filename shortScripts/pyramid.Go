package main

import "fmt"

func main()  {

	for i := 1; i < 6; i++{
		//row := ""
		//the row method also works here
		for j := 1; j < i; j++ {
			fmt.Print("#.")
		}
		fmt.Println()
	}

}