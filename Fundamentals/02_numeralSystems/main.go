package main

import "fmt"

//%d - decimal, %b - binary, %x - hexadecimal using Printf - print format
func main() {
	for i := 60; i < 200; i++ {
		fmt.Printf("%d - %b - %x - %q \n", i, i, i, i)
	}

}
