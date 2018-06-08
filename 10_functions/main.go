package main

import "fmt"

func main() {

	greet("Jane")
	greet("John")
	greet2("Jane", "Doe")
	fmt.Println(sprint("Jane ", "Doe"))
	fmt.Println(multisprint("Jane ", "Doe "))
}

func greet(name string) { //single argument
	fmt.Println(name)
}

func greet2(fname, lname string) { //multiple argument, no return
	fmt.Println(fname, lname)
}

func sprint(fname, lname string) string { //multiple arguments, single return
	return fmt.Sprint("Welcome ", fname, lname)
}

func multisprint(fname, lname string) (string, string) { //multiple arguments and returns
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
