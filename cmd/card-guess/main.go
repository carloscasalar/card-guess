package main

import (
	"fmt"
	"os"

	"github.com/carloscasalar/card-guess/internal/deck"
)

const TrickSampleSize = 21

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	dealer := deck.NewDealer(1)

	dealer.ShuffleCards()

	cards := make([]deck.Card, TrickSampleSize)
	for i := 0; i < 21; i++ {
		card, err := dealer.Deal()
		if err != nil {
			return fmt.Errorf("unexpected error while dealing the card %vth: %w", i+1, err)
		}
		cards[i] = card
	}

	fmt.Println(cards)
	fmt.Println("Pick a card from above and hold it in your mind, then press enter.")

	return nil
}
