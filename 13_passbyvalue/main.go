package main

import "fmt"

func main() {

	age := 44

	fmt.Println(age)	//44
	fmt.Println(&age) //0xc420014058

	changeMe(&age)

	fmt.Println(age)	//24

}

func changeMe(z *int) {
	fmt.Println(z)		//0xc420014058
	fmt.Println(*z)		// 44

	*z = 24
}
