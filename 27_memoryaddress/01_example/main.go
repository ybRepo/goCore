package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x) //x starts as 5
	zero(x)
	addr(x)

}

func zero(x int) { //x is passed to func zero

	x++
	fmt.Println(x) //x is now 6
}

func addr(x int) { //x is passed to func addr

	var y = &x      //var y is pointing to x's mem address
	fmt.Println(y)  //printing y's value = x's mem address
	fmt.Println(&y) //printing y's mem address which is different than x's

	*y = 11        //set the value at y's pointer (x's address) to 11
	fmt.Println(x) //demonstrate that values can change through mem address

}
