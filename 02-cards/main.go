package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand: ", hand)
	fmt.Println("Remaining Deck : ", remainingDeck)
	cards.saveToFile("my_cards")
}
