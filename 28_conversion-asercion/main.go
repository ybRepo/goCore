package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12"
	var y = 6

	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)

	var name interface{} = "sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)

	} else {
		fmt.Printf("Value is not a string\n")
	}
}
