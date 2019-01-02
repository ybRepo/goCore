package main

import "fmt"

type person struct { //a new type person which contains the following fields
	first string
	last  string
	age   int
}

func (p person) fullname() string { //a new function the Receiver , func name, the parameters, the return, the body
	return p.first + p.last //read as follows: any value (p) of type "person" can access fullname method. This method returns a string
} //consisting of p.first and p.last

func main() {

	p1 := person{"James ", "Bond", 20}      //create the first person of type person in value p1
	p2 := person{"Miss ", "Moneypenny", 18} //create the second person of type person in value p2
	fmt.Println(p1.fullname())              //for the first person (p1) call the function "fullname" and print the return
	fmt.Println(p2.fullname())              //for the second person (p2) call the function "fullname" and print the return

}
