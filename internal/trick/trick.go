package trick

import (
	"fmt"

	"github.com/carloscasalar/card-guess/pkg/threepilestrick"

	"github.com/carloscasalar/card-guess/internal/mat"

	"github.com/carloscasalar/card-guess/internal/deck"
)

const trickSampleSize = 21

type Trick interface {
	Sample() threepilestrick.Pile
	Mat() mat.Mat
}

func New(shuffleBeforeInitialDraw bool) (Trick, error) {
	dealer := deck.NewDealer()
	if shuffleBeforeInitialDraw {
		dealer.ShuffleCards()
	}
	cards := make([]threepilestrick.Card, trickSampleSize)
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
	sample threepilestrick.Pile
	mat    mat.Mat
}

func (s startingTrick) Sample() threepilestrick.Pile {
	return s.sample
}

func (s startingTrick) Mat() mat.Mat {
	return s.mat
}
