package packages

import (
	"math/rand"
	"time"
)

var deck = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}	//need to fix this ...

//GetCard random card from deck
func GetCard() int {

	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := r.Intn(len(deck))

	return deck[i]

}
