package main

import "fmt"

func main() {

	//initializing a variable is when you give it its type and value
	a := 10
	b := "golang"
	c := 4.17
	d := true
	name := "Yashar"

	//variable with zero value
	var e bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Println("Hello", name, "!")
}
