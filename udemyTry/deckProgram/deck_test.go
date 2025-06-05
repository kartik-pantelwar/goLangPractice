package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d1 := NewDeck()
	if len(d1) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d1))
	}

	if d1[0] != "club-2" {
		t.Errorf("Expected the first Card to be club-2, but got %v", d1[0])
	}
	if d1[len(d1)-1] != "spade-ace" {
		t.Errorf("Expected the last Card to be spade-ace, but got %v", d1[len(d1)-1])
	}
}

func TestSaveToFileAndnewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d1:=NewDeck()
	d1.saveToFile("_decktesting")

	d2:=newDeckFromFile("_decktesting.txt")
	if len(d2)!=52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d2))
	}
	os.Remove("_decktesting.txt")
}