package main

import (
	"fmt"
	"sort"
)

//2 examples of how to sort a slice of string.
//The first method looks at creating a new type which has functions equivalent tot he Interface interface type. This is because the sort package has a method sort.Sort which can take an Interface.
//So the type people is given the same functions as the Interface function.
type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(studyGroup) //therefore when a slice of people is added to a variable, the sort.Sort method can be used.
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))


	fmt.Println(studyGroup)

	useinterface() //a much easier way to sort the list is to call the sort.Strings method
}

func useinterface() {
	studyGroup2 := []string{"Allan", "Zack", "Ryan", "Todd"}
	studyGroup3 := []string{"Allan", "Shawn", "Ryan", "Michael"}

	sort.Strings(studyGroup2)
	sort.StringSlice(studyGroup3).Sort()
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup3)))

	fmt.Println(studyGroup2)
	fmt.Println(studyGroup3)
}
