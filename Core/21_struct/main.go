package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type person2 struct {
	person        //inherits struct
	sex    string //adds new fields
}

func main() {

	p1 := person{"James", "Bond", 20} //structs can have their values initialized in the same order as they are defined in the struct
	p2 := person{"Miss", "Monteray", 18}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2)

	p3 := person2{
		person: person{ //for embedded struct, a struct litterals is better suited as you initialize values explicitly .
			first: "Samantha", //for each field rather than worrying about entering values in order of the struct
			last:  "Knox",
			age:   21,
		},
		sex: "Female",
	}

	fmt.Println(p3.person.first, p3.last, p3.age, p3.sex) //p3 values can be returned by going through person2 struct like so "p3.sex" or by drilling down "p3.person.first"
}
