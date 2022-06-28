package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if !(len(d) <= 52) {
		t.Errorf("Too many cards. Expected deck length <= 52, got %v", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	tp := "_decktesting"
	os.Remove(tp)
	deck := newDeck()
	deck.saveToFile(tp)

	loadedDeck := newDeckFromFile(tp)

	if !(len(loadedDeck) <= 52) {
		t.Errorf("Too many cards. Expected deck length <= 52, got %v", len(loadedDeck))
	}

	os.Remove(tp)
}
