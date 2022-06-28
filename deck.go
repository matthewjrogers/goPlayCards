package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type 'deck'
// a slice of strings
type deck []string

// generate a new deck
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

// print all cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal a hand of cards
func deal(d deck, nCards int) (deck, deck) {

	return d[:nCards], d[nCards:]
}

// convert deck to a string to facilitate saving
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// function to write deck to disk
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// create a new deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	for i := range d {
		elem := rand.Intn(len(d) - 1)
		d[i], d[elem] = d[elem], d[i]
	}
}
