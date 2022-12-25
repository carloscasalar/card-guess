package main

import (
	"fmt"
	"github.com/carloscasalar/card-guess/internal/deck"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	dealer := deck.NewDealer(1)

	dealer.ShuffleCards()

	cards := make([]deck.Card, 21)
	for i := 0; i < 21; i++ {
		card, err := dealer.Deal()
		if err != nil {
			fmt.Printf("Unexpected error while dealing the card %vth", i+1)
			os.Exit(1)
		}
		cards[i] = card
	}

	fmt.Println(cards)
	fmt.Println("Pick a card from above and hold it in your mind, then press enter.")

	return nil
}
