package main

import (
	"fmt"

	"gophercises/deck"
)

func main() {
	d := deck.New("", 3, []string{"2", "3", "4", "5", "6", "7", "8"})
	for _, card := range d {
		fmt.Print(card.Rank, string(card.Suit), " ")
	}
	fmt.Println()
}
