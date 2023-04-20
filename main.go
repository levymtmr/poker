package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Value string
	Suit  string
}

func NewDeck() []Card {
	var deck []Card

	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}

	for _, suit := range suits {
		for _, value := range values {
			card := Card{value, suit}
			deck = append(deck, card)
		}
	}
	return deck
}

func ShuffleDeck(deck []Card) []Card {
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func main() {
	cards := NewDeck()
	fmt.Println(cards)
	fmt.Println(ShuffleDeck(cards))

}
