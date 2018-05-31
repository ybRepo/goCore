package main

import "fmt"

func main() {
	name := make([]string, 4)

	name[0] = "Mehdi"
	name[1] = "Tim"
	name[2] = "Jeremy"
	name[3] = "Trevor"

	for i := 0; i < len(name); i++ {

		fmt.Println("Is", name[i], "in the classroom?")

		switch name[i] {

		case name[0]:
			fmt.Println(name[i], "Here")

		case name[1]:
			fmt.Println(name[i], "Here")

		case name[2]:
			fmt.Println(name[i], "Here")

		case name[3]:
			fmt.Println(name[i], "Here")

		}
		if i == len(name[1:]) {

			fmt.Println("All", len(name), "students are accounted for, welcome to Golang!")
		}
	}

}
