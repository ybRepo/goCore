package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}
	nested()
	condloop()
	breakloop()
	continueloop()
}

func nested() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}
}

func condloop() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func breakloop() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}

func continueloop() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue	//continue loops through the condition until a value is false
		}
		fmt.Println(i)
		if i >= 10 {
			break
		}
	}

}
