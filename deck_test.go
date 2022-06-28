package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if !(len(d) <= 52) {
		t.Error("Too many cards. Expected deck length <= 52, got", len(d))
	}
}
