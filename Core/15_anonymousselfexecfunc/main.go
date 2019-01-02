package main

import "fmt"

func main() {
	func() {
		fmt.Println("i'm driving!")
	}() //the () executes the nameless function and prints i'm driving.
}
