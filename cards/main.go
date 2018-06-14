package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("test.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
