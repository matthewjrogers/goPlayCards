package main

import (
	"fmt"
	"strings"
)

// create a new type 'deck'
// a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	ranks := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, rank+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, nCards int) (deck, deck) {

	return d[:nCards], d[nCards:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
