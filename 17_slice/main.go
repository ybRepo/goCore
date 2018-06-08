package main

import "fmt"

func main() {
	shorthandslice()
	makeslice()

}

func shorthandslice() {
	myslice := []int{1, 2, 3, 4, 5, 9, 11}
	fmt.Printf("%T\n", myslice)
	fmt.Println(myslice)
	fmt.Println(myslice[2:4])
	fmt.Println(myslice[2])
	fmt.Println("mySlice"[2])
}

func makeslice() {
	records := make([]string, 3, 6)

	records[0] = "Monday"
	records[1] = "Tuesday"
	records[2] = "Wednesday"
	records = append(records, "Thursday")
	records = append(records, "Friday")
	records = append(records, "Saturday")
	records = append(records, "Sunday")

	fmt.Println(records[0:6])
	fmt.Println(len(records), cap(records))
}
