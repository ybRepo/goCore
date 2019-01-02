package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) //empty slice
	changeMe(m)
	fmt.Println(m) //can be seen to impact m as a slice is a reference
	// type which has the address explicitly in it.

}

func changeMe(z []string) {
	z[0] = "todd"  //passed slice to z and set value in position 0 to todd
	fmt.Println(z) //position 0 prints "todd"

}
