package core

import (
	"math/rand"

	"github.com/levymtmr/poker/models"
)

func NewDeck() []*models.Card {
	var deck []*models.Card

	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}

	for _, suit := range suits {
		for _, value := range values {
			card := &models.Card{Value: value, Suit: suit}
			deck = append(deck, card)
		}
	}
	return deck
}

func ShuffleDeck(deck []*models.Card) []*models.Card {
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}
