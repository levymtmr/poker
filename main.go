package main

import (
	"fmt"

	"github.com/levymtmr/poker/core"
)

func main() {
	cards := core.NewDeck()
	shuffledCards := core.ShuffleDeck(cards)

	for i := 0; i < len(shuffledCards); i++ {
		fmt.Println(*shuffledCards[i])
	}

	// fmt.Println(cards)
	// fmt.Println(core.ShuffleDeck(cards))

}
