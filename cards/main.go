package main

import (
	"fmt"
)

func main() {
	//var card string = "Ace of Spades"
	firstCard := "Ace of Spades"
	fmt.Println(firstCard)

	//cards := []string{"Two of Spades", newCard(), "Seven of Hearts"}
	cards := deck{card{value: "5", suit: "S"}, card{value: "A", suit: "H"}}

	cards = append(cards, card{value: "9", suit: "D"})
	fmt.Println(cards)

	cards.print()

	newDeck := newDeck()
	newDeck.print()
	//newDeck.writeToFile("newDeckFile.txt")

	fileDeck := newDeckFromFile("deckfile.txt")
	fmt.Println("from file:")
	fileDeck.print()

	fileDeck.shuffle()

	fileDeck.print()

}
