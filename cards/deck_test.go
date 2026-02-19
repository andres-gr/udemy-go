package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	l := len(d) - 1

	if d[l] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[l])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
