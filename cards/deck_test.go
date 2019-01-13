package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d := newDeck()
	length := 16

	if len(d) != length {
		t.Errorf("Expected a deck of length %v but got %v", length, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of \"Ace of spades\" but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of \"Four of Clubs\" but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	length := 16
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != length {
		t.Errorf("Expected a deck of length %v but got %v", length, len(loadedDeck))
	}

	os.Remove("_decktesting")


}