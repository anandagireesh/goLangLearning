package main

import "fmt"

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {

		for _, value := range cardValues {

			cards = append(cards, value+" Of "+suit)

		}

	}

	return cards
}

func (d deck) print() {

	for i, card := range d {
		var j int = i + 1
		fmt.Println(j, card)

	}

	//fmt.Println("my card = " + d[1])

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
