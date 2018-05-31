package main

import "fmt"

const (
	a = 1
	b = "this is a constant"
	c = true
)

const z string = "a typed constant, while the previous were untype constants since their types were not defined"

func main() {
	const x = "death & taxes"
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(true)
	fmt.Println(z)
}
