package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.print()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand: ", hand)
	fmt.Println("Remaining Deck : ", remainingDeck)
}
