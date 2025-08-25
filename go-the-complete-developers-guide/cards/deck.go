package main

import (
	"fmt"
	"math/rand"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// function that has a receiver inside the parens before the function name
// basically, what ever is calling this function is passed into the function itself
// kind of like a 'this' or 'self' in other languages
// also, it is a convention to use the first letter as the variable name for the receiver object
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// function that takes 2 parameters
// the first parameter is the deck to deal from
// the second parameter is the number of cards to deal
// the function returns 2 decks -- the (deck, deck) part
func deal(d deck, handSize int) (deck, deck) {
	hand := deck{}
	deckCopy := make(deck, len(d))
	copy(deckCopy, d)

	for i := 0; i < handSize; i++ {
		card, remainingDeck := deckCopy.pickRandomCard()
		hand = append(hand, card)
		deckCopy = remainingDeck
	}

	return hand, d
}

// function that takes a deck as a parameter
// the function returns a string and a deck
// the string is the card that was picked
// the deck is the remaining deck
func (d deck) pickRandomCard() (string, deck) {
	randomIndex := rand.Intn(len(d))
	card := d[randomIndex]

	// this effectively removes the card from the deck
	d[randomIndex] = d[len(d)-1]
	d = d[:len(d)-1]

	return card, d
}
