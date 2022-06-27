package main

func main() {
	cards := newDeck()

	hand, cards := deal(cards, 3)
	hand.print()
	cards.print()
}
