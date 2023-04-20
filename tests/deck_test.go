package tests

import (
	"reflect"
	"testing"

	"github.com/levymtmr/poker/core"
)

func TestNewDeck(t *testing.T) {
	deck := core.NewDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	firstCard := deck[0]
	if firstCard.Value != "Ace" {
		t.Errorf("Expected first card value of Ace, but got %v", firstCard.Value)
	}

	lastCard := deck[len(deck)-1]
	if lastCard.Suit != "Spades" {
		t.Errorf("Expected last card suit of spades, but got %v", lastCard.Value)
	}
}

func TestShuffleDeck(t *testing.T) {
	deck := core.NewDeck()
	shuffledDeck := core.ShuffleDeck(deck)

	if reflect.DeepEqual(deck, shuffledDeck) {
		t.Errorf("Expected deck to be shuffled, but it's the same as the unshuffled deck")
	}
}
