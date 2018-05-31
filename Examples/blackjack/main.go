package main

import (
	"fmt"
)

func main() {

	//firstcard := dealhand()
	//fmt.Println(firstcard)
	getCard()

}

func getCard() {
	cards := [][]string{}

	// These are the first two rows.
	row1 := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	row2 := []string{"Club", "Spade", "Diamond", "Heart"}

	// Append each row to the two-dimensional slice.
	cards = append(cards, row1)
	cards = append(cards, row2)



	// Access an element.
	fmt.Println("First element")
	fmt.Println(cards[0][1])

	// Display entire slice.
	fmt.Println("Values")
	fmt.Println(cards)
}

