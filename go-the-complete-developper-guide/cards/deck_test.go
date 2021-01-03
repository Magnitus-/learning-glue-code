package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected \"Two of Spades\" as first value, got %q", d[0])
	}

	if d[len(d) -1] != "Ace of Clubs" {
		t.Errorf("Expected \"Ace of Clubs\" as first value, got %q", d[len(d) -1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	cards := newDeck()
	cards.saveToFile("_decktesting")
	
	newCards := newDeckFromFile("_decktesting")

	if len(newCards) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(newCards))
	}

	if newCards[0] != "Two of Spades" {
		t.Errorf("Expected \"Two of Spades\" as first value, got %q", newCards[0])
	}

	if newCards[len(newCards) -1] != "Ace of Clubs" {
		t.Errorf("Expected \"Ace of Clubs\" as first value, got %q", newCards[len(newCards) -1])
	}

	os.Remove("_decktesting")
}