package main

import "fmt"

func main() {
	greatest := max(4, 7, 9, 113, 545, 76, 48, 39)
	fmt.Println(greatest)

}

func max(numbers ...int ) int {
	var largest int

	for _, v :=range numbers{
		if v > largest{
			largest = v
		}
	}
	return largest
}