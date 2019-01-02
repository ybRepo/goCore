package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x + 1)
		y := "44"
		fmt.Println(y)
	}
	fmt.Println(x)

	fmt.Println("\n***********\n" +
		"While X is declared within the main func it is accessible within inner scope of {} and can be manupilated X+1,\n" +
		"While Y is declared within the inner scope {} it cannot be called from outside of {}\n" +
		"X is called again to demonstrate the same value since the changes to X were limited to inner scope {}" +
		"\n***********\n")
}
