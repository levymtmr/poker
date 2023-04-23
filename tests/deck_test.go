package tests

import (
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
	for i := 0; i < 10; i++ {
		shuffledDeck := core.ShuffleDeck(deck)

		if len(shuffledDeck) != len(deck) {
			t.Errorf("Expected both deck has the same size")
		}

		for i, card := range shuffledDeck {
			if card.Value != deck[i].Value || card.Suit != deck[i].Suit {
				t.Errorf("Expected difference between value position")
			}
		}

	}
}
