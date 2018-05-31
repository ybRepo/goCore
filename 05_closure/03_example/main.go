package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("\n***********\n" +
		"To simplify the code from example 2, the main function contains an inner function\n" +
		"However, the inner function is assigned to a variable and takes x and adds +1\n" +
		"Note that the inner function has no identifier func() a.k.a. an anonymous function\n" +
		"as opposed to func name () and the increment variable is called twice in a row" +
		"\n***********\n")
}
