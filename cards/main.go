package main

func main(){
	cards := newDeckFromFile("somegofile")

	//cards.print()

	//hand, remainingDeck := deal(cards, 5)

	//hand.print()
	//remainingDeck.print()

	//println(cards.toString())
	//cards.saveToFile("somegofile")
	cards.shuffle()
	cards.print()

}

