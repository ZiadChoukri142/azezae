package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	deckOne := "1, 2, 3"
	deckTwo := "4, 5, 6"
	deckThree := "7, 8, 9"
	deckFour := "10, 11, 12"
	fmt.Printf("Player 1: %v\n", deckOne)
	fmt.Printf("Player 2: %v\n", deckTwo)
	fmt.Printf("Player 3: %v\n", deckThree)
	fmt.Printf("Player 4: %v\n", deckFour)
}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
