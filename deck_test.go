package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected deck Four of Clubs, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting.txt")

	deck := newDeck()
	deck.saveToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")
}