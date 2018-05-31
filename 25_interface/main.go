package main

import (
	"fmt"
	"math"
)

type Shape interface { //since the methods for Circle and Square have the same signature "area() float64" a new type Shape of type interface can be created.
	area() float64 //this interface calls the area function and returns float64
	//interfaces are types that declare behaviour. The behaviour is never implemented by the inferece type directly, but instead by user defined types via methods.
}

type Square struct { //create a new type called square with field side
	side float64
}

type Circle struct { //create new type called circle with field radius
	radius float64
}

func (s Square) area() float64 { //function for Square which returns a float by multiplying fields side^2
	return s.side * s.side
}

func (c Circle) area() float64 { //function for circle which returns a float by multiplying Pi*radius^2
	return math.Pi * c.radius * c.radius
}

func measure(z Shape) { //new function which takes a Shape interface, and prints the results from area function
	fmt.Println(z.area())
}

func main() {
	s := Square{10} //set the field side of a Square strut to 10
	measure(s)      //use the interface to pass the values to measure function which prints the area using the appropriate (square) area () function

	c := Circle{20}
	measure(c)
}
