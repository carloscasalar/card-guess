package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"

	"github.com/carloscasalar/card-guess/internal/trick"

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

	sample := deck.NewPile()
	for i := 0; i < TrickSampleSize; i++ {
		card, err := dealer.Deal()
		if err != nil {
			return fmt.Errorf("unexpected error while dealing the card %vth: %w", i+1, err)
		}
		sample = sample.AddCard(card)
	}

	fmt.Println(sample.String())
	fmt.Println("Pick a card from above and hold it in your mind.")
	fmt.Println("Now I'll split the cards into three piles, watch your card.")
	fmt.Println("Then press enter to continue")
	_, _ = fmt.Scanln()

	mat := trick.NewMat()
	for {
		var card deck.Card
		var err error
		card, sample, err = sample.DrawCard()
		if err != nil {
			if errors.Is(err, deck.ErrNoMoreCardsInThePile) {
				break
			}
			return err
		}
		mat.PlaceIntoNextPile(card)
	}

	prompt := promptui.Select{
		Label: "Select the pile where your card is",
		Items: mat.Piles(),
	}

	_, result, err := prompt.Run()

	if err != nil {
		return fmt.Errorf("prompt failed %w", err)
	}

	fmt.Printf("You choose %q\n", result)

	return nil
}
