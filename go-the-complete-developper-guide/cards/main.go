package main

func main() {
	//cards := newDeck()
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()
	cards := newDeck()
	cards.saveToFile("cards.txt")
	newCards := newDeckFromFile("cards.txt")
	newCards.shuffle()
	newCards.print()
}