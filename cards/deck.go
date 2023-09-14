package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type "deck"

type deck []card

func newDeck() deck {
	cards := deck{} //assign to empty deck

	suits := []string{"S", "H", "C", "D"} //Spades, Hearts, Clubs, Diamonds
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, v := range values {
		for _, s := range suits {
			cards = append(cards, card{suit: s, value: v})
		}
	}
	return cards
}

func newDeckFromFile(filename string) deck {

	tempBytes, error := os.ReadFile(filename)

	if error != nil {
		fmt.Println(filename, "does not exist. Creating from scratch.")
		return newDeck()
	} else {
		ts := strings.Split(string(tempBytes), ",") //get string slice

		dd := deck{}

		for _, tb := range ts {
			if len(tb) == 2 {
				dd = append(dd, card{value: tb[:1], suit: tb[1:]})
			} else {
				dd = append(dd, card{value: tb[:2], suit: tb[2:]})
			}

		}

		return dd
	}
}

func (d deck) shuffle() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)
	l := len(d) - 1

	for i := range d {
		r_index := r.Intn(l)
		d[i], d[r_index] = d[r_index], d[i]
	}
}

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, printCard(c))
	}
}

func deal(d deck, handSize int) (deck, deck) {
	//1st deck is hand, 2nd deck is remaining deck of cards
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	var ds []string

	for _, c := range d {
		ds = append(ds, printCard(c))
	}

	return strings.Join([]string(ds), ",")
}

func (d deck) writeToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0664)

}
