package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)

}
