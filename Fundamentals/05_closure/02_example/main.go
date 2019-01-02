package main

import "fmt"

var x = 0

func increment() int {

	x++
	return x

}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("\n***********\n" +
		"Variable X is declared at the package level and initalized with a value of 0\n" +
		"the increment func return an int by taking the value X and adding +1\n" +
		"main func calls the increment function twice, thus impacting the value of X each time" +
		"\n***********\n")
}
