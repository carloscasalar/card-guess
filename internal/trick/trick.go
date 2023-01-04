package trick

import (
	"fmt"

	"github.com/carloscasalar/card-guess/internal/deck"
)

const trickSampleSize = 21

type Trick interface {
	Cards() []deck.Card
}

type startingTrick struct {
	cards []deck.Card
}

func (s startingTrick) Cards() []deck.Card {
	return s.cards
}

func New() (Trick, error) {
	dealer := deck.NewDealer()
	cards := make([]deck.Card, trickSampleSize)
	for i := 0; i < trickSampleSize; i++ {
		card, err := dealer.Deal()
		if err != nil {
			return nil, fmt.Errorf("unexpected error while dealing the card %vth: %w", i+1, err)
		}
		cards[i] = card
	}

	return &startingTrick{cards}, nil
}
