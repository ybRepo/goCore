package main

import (
	"fmt"
	"goTraining/04_scope/visibility"
)

//the variables declared at the package level can be out of order and therefore appear at the end.
var x = 42
var z = 1

//main function that prints the variable x within package main scope
func main() {
	fmt.Println(x, "- package variable x")

	//variables declared within a function must adhere to an order
	y := 17
	fmt.Println(y, "-function vairable y")

	//since program runs function main, you need to call other functions if needed like function foo within main
	foo()

	//calls a variable declared in the visibility package that is visible to the rest of the program
	fmt.Println(visibility.MyName)

	//calls a function declared in the visibility package that is visible to the rest of the program
	visibility.PrintVar()

}

func foo() {
	fmt.Println(x + z)
}
