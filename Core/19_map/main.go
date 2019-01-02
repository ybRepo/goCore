package main

import "fmt"

var myGreeting = make(map[string]string) //create a map of key pairs of type string

func main() {
	addmap()
	fmt.Println(myGreeting)

	updatemap()
	fmt.Println(myGreeting)

	deletemapkey()
	fmt.Println(myGreeting)

	litteralmap()

	keypairexists()

}

func addmap() {
	myGreeting["Tim"] = "Good Morning!" //add new key pair value to a map
	myGreeting["Jeremy"] = "Hello!"
	myGreeting["Harleen"] = "Howdy!"

}

func updatemap() {
	myGreeting["Harleen"] = "Bonjourno!" //the key "Harleen" must be exact, and of the same case. "harleen" would add a new record instead of updating "Harleen"
}

func deletemapkey() {
	delete(myGreeting, "Harleen")
}

func litteralmap() { //this is the alternate way of creating maps where values are defined
	myOtherGreetings := map[string]string{
		"Frank": "Good Day Sir!",
		"Jenny": "Hi!",
	}
	fmt.Println(myOtherGreetings)
}

//This is accomplished by if statements in Go can include both a condition and an initialization statement. The example below uses both:
// 			initializes two variables - val will receive either the value of "foo" from the map or a "zero value" (in this case the empty string) and ok will receive
// 										a bool that will be set to true if "foo" was actually present in the map evaluates ok, which will be true if "foo" was in the map
//			If "foo" is indeed present in the map, the body of the if statement will be executed and val will be local to that scope.

func keypairexists() {
	if val, ok := myGreeting["Harleen"]; ok {
		fmt.Println("exists:", ok)
		fmt.Println("val:", val)

	} else {
		fmt.Println("exists:", ok)
		fmt.Println("The key doesn't exist.")

	}

}
