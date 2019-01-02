package main

import "fmt"

func main() {
	a := 32

	fmt.Println(a)  //print the value of a
	fmt.Println(&a) //print the memory address storing a

	var b = &a // new var b is pointing* to an int in memory location &a

	fmt.Println(b)  //print the value of b - which is the pointer
	fmt.Println(*b) //print the value at b's address by dereferencing

}
