package packages

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var dealertotal = 0
var playertotal = 0
var playerhand []int
var dealerhand []int

//Play called by main to start the game
func Play() {
	dealStartingHand()

	if bj := checkBlackjack(); bj != true {

		if pb := playerTurn(); pb != true {

			if db := dealerTurn(); db != true {
				winner()
			}
		}
	}
}

func handValue(x []int) int {
	total := 0
	ace := 0

	for _, v := range x {
		if v == 1 {
			ace++
		} else {
			total += v
		}
	}

	switch ace {

	case 1:
		if total <= 10 {
			total += 11
		} else {
			total ++
		}

	case 2:
		if total <= 8 {
			total += 12
		} else {
			total += 2
		}

	case 3:
		if total <= 7 {
			total += 13
		} else {
			total += 3
		}

	case 4:
		if total <= 6 {
			total += 14
		} else {
			total += 4
		}
	}

	return total
}

func checkBlackjack() bool {
	blackjack := false
	if playertotal == 21 && dealertotal == 21 {
		fmt.Println("Draw!")
		blackjack = true

	} else if playertotal == 21 {
		fmt.Println("Blackjack, you WIN,")
		blackjack = true

	} else if dealertotal == 21 {
		fmt.Println("You Lose! \n")
		blackjack = true

	}
	return blackjack

}

func dealStartingHand() {

	for i:=0; i<2; i++{
		time.Sleep(500 * time.Millisecond)
		dealerhand = append(dealerhand, GetCard())

		time.Sleep(500 * time.Millisecond)
		playerhand = append(playerhand, GetCard())
	}

	dealertotal = handValue(dealerhand)
	playertotal = handValue(playerhand)

	fmt.Println("Dealer has:", dealerhand[0])
	fmt.Println(dealerhand[0], ", ?")

	fmt.Println("You have:", playertotal)
	fmt.Println(playerhand)
}

func playerTurn() bool {
	bust := false
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; playertotal <= 21; i++ {
		fmt.Print("Would you like another card? ")
		scanner.Scan()
		response := scanner.Text()

		if response == "y" {
			playerhand = append(playerhand, GetCard())
			playertotal = handValue(playerhand)

			fmt.Println(playerhand)
			fmt.Println("You have:", playertotal)

			if playertotal > 21 {
				bust = true
				fmt.Println("Bust! You lose.")
			}

		} else {
			fmt.Println("Player chose to stay on:", playertotal)
			break
		}
	}

	return bust
}

func dealerTurn() bool {
	bust := false

	time.Sleep(300 * time.Millisecond)
	fmt.Println("Dealer shows: ", dealertotal)
	fmt.Println(dealerhand)

	time.Sleep(300 * time.Millisecond)

	for i := 0; dealertotal <= 18; i++ {

		time.Sleep(1000 * time.Millisecond)
		if dealertotal == 18 {
			fmt.Println("Dealer stays on 18")
			break
		}

		if dealertotal < 18 {
			dealerhand = append(dealerhand, GetCard())
			dealertotal = handValue(dealerhand)
			fmt.Println("Dealer hits:")
			fmt.Println(dealerhand, "Dealer has:", dealertotal)
		}
	}

	if dealertotal > 21 {
		bust = true
		fmt.Println("Dealer Busts! - You Win!")

	}

	return bust

}

func winner() {
	if dealertotal == playertotal {
		fmt.Println("Draw!")
	} else if dealertotal < playertotal {
		fmt.Println("You WIN!")
	} else {
		fmt.Println("Dealer WINS!")
	}

}
