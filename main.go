package main

import "fmt"

func main() {

	//var card string = "Ace of Spades"
	//card := newCard()

	// cards = append(cards, newCard())

	cards := newDeck()

	// to string

	//fmt.Println(cards.toString())

	fmt.Println("My cards from file")
	myCards := newDeckFromFile("mycard")
	myCards.print()
	fmt.Println(".......................................................................................")
	fmt.Println(".......................................................................................")
	fmt.Println(".......................................................................................")
	fmt.Println(".......................................................................................")

	cards.shuffle()
	cards.print()
	//cards.saveToFile("mycard")

	// hand, remainingCards := deal(cards, 6)

	// // cards.print()
	// fmt.Println("Cards in Hand")
	// fmt.Println("--------------")

	// hand.print()

	// fmt.Println("Remaining Cards")
	// fmt.Println("------------------")
	// remainingCards.print()
}

// func newCard() string {

// 	return "Two of Spades"

// }
