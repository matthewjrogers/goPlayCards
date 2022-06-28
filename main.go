package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	//hand, cards := deal(cards, 3)
	//hand.print()
	cards.print()
	//cards.saveToFile("my_cards")
}
