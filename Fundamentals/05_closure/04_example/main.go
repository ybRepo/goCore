package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x

	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("\n***********\n" +
		"The main func has a variable called increment which is assigned a function called wrapper\n" +
		"func wrapper returns an anonymous func. This anonymous function returns an int\n" +
		"so x is initialized to 0, and the anonymized func is called at time of calling return\n" +
		"X is added +1 and wrapper returns the value x from func() int" +
		"\n***********\n")
}
