package main

import "fmt"

func main() {
	cards := newDeck()

	hand1, remainingDeck := deal(cards, 5)
	hand2, remainingDeck := deal(remainingDeck, 5)

	cards = remainingDeck
	hand1.print()
	fmt.Println(hand1.toString())
	if hand1.saveToFile("hand1.txt") != nil {
		fmt.Println("Error saving hand1 to file")
	} else {
		fmt.Println("Hand1 saved to file")
	}
	fmt.Println("&&&&&&&&&&")
	hand2.print()
	fmt.Println(hand2.toString())
	if hand2.saveToFile("hand2.txt") != nil {
		fmt.Println("Error saving hand2 to file")
	} else {
		fmt.Println("Hand2 saved to file")
	}

	fmt.Println("Loading hand1 from file")
	loadedHand1 := deck{}.loadFromFile("hand1.txt")
	loadedHand1.print()

	cards2 := newDeck()
	fmt.Println("\n\n\nShuffling cards2")
	fmt.Println("Before shuffle:")
	fmt.Println(cards2.toString())
	cards2.shuffle()
	fmt.Println("After shuffle:")
	fmt.Println(cards2.toString())
}
