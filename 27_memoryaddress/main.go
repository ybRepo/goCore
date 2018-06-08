package main

import "fmt"

func main() {
	a := 42

	fmt.Println("a - ", a)
	fmt.Println("a - ", &a)
	swim()
}

const metersToYards = 1.09361

func swim() {
	var meters float64
	fmt.Print("Enter meteres swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "meter is", yards, "yards.")

}
