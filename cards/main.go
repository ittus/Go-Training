package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("test.txt")
	cards := newDeckFromFile("test.txt1")
	cards.print()
}
