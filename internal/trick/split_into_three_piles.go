package trick

import (
	"errors"

	"github.com/carloscasalar/card-guess/pkg/threepilestrick"

	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/mat"
)

func splitIntoThreePiles(sample threepilestrick.Pile) (*threepilestrick.Mat, error) {
	theMat := mat.New()
	for {
		var card threepilestrick.Card
		var err error
		card, sample, err = sample.DrawCard()
		if err != nil {
			if errors.Is(err, deck.ErrNoMoreCardsInThePile) {
				break
			}
			return nil, err
		}
		theMat = theMat.PlaceIntoNextPile(card)
	}
	return &theMat, nil
}
