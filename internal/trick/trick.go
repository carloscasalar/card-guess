package trick

import (
	"fmt"

	"github.com/carloscasalar/card-guess/internal/mat"

	"github.com/carloscasalar/card-guess/internal/deck"
)

type Trick interface {
	Sample() deck.Pile
	Mat() mat.Mat
}

const trickSampleSize = 21

func New(shuffleBeforeInitialDraw bool) (Trick, error) {
	dealer := deck.NewDealer()
	if shuffleBeforeInitialDraw {
		dealer.ShuffleCards()
	}
	cards := make([]deck.Card, trickSampleSize)
	for i := 0; i < trickSampleSize; i++ {
		card, err := dealer.Deal()
		if err != nil {
			return nil, fmt.Errorf("unexpected error while dealing the card %vth: %w", i+1, err)
		}
		cards[i] = card
	}

	sample := deck.NewPile(cards...)
	theMat, err := splitIntoThreePiles(sample)
	if err != nil {
		return nil, err
	}

	return &startingTrick{sample, *theMat}, nil
}

type startingTrick struct {
	sample deck.Pile
	mat    mat.Mat
}

func (s startingTrick) Sample() deck.Pile {
	return s.sample
}

func (s startingTrick) Mat() mat.Mat {
	return s.mat
}
