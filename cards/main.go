package main

func main() {
	cards := newDeck()
	cards.shuffle()

	// hand, remainingDeck := deal(cards, 5)
	//
	// hand.print()
	//
	// fmt.Println()
	// fmt.Println("------------")
	// fmt.Println()
	//
	// remainingDeck.print()

	// err := cards.saveToFile("cards_file")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cards.print()
}
