package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := len(d)

	if l != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if printCard(d[0]) != "AS" {
		t.Errorf("Expected first card is Ace of Spades, but got %v", printCard(d[0]))
	}
	if printCard(d[l-1]) != "KD" {
		t.Errorf("Expected last card is King of Diamonds, but got %v", printCard(d[l-1]))
	}

}

func TestNewDeckFromFile(t *testing.T) {
	testfilename := "_decktesting"
	os.Remove(testfilename) //cleanup environment

	d := newDeck()
	d.writeToFile(testfilename)

	fileDeck := newDeckFromFile(testfilename)

	l := len(fileDeck)

	if l != 52 {
		t.Errorf("Expected 52 card but found %v", l)
	}

}
