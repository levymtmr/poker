package main

import (
	"fmt"

	"github.com/levymtmr/poker/core"
)

func main() {
	cards := core.NewDeck()
	fmt.Println(cards)
	fmt.Println(core.ShuffleDeck(cards))
}
