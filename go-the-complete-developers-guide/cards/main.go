package main

import "fmt"

func main() {
	cards := newDeck()

	hand1, remainingDeck := deal(cards, 5)
	hand2, remainingDeck := deal(remainingDeck, 5)

	cards = remainingDeck
	hand1.print()
	fmt.Println("&&&&&&&&&&")
	hand2.print()
}
