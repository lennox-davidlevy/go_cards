package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		//%v injects value in to string, otherwise error
		t.Errorf("Expected length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element to be 'Ace of Spades', but got %v", d[0])
	}
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last element to be 'King of Diamonds' but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")
	if d.toString() != loadedDeck.toString() {
		t.Errorf("Expected file saved to be same as file retrieved")
	}
	os.Remove("_deckTesting")
}
