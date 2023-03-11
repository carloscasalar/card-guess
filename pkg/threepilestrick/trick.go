package threepilestrick

import (
	"errors"
	"fmt"

	"github.com/carloscasalar/card-guess/internal/mat"

	"github.com/carloscasalar/card-guess/internal/deck"
)

type Card string
type Pile []Card

type Trick interface {
	Sample() Pile
	FirstPile() Pile
	SecondPile() Pile
	ThirdPile() Pile
	GuessMyCard() (*Card, error)
	MyCardIsInPile(PileHolder) (Trick, error)
}

func New(shuffleBeforeInitialDraw bool) (Trick, error) {
	const trickSampleSize = 21
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
	return intermediateTrickState{
		sample:          sample,
		mat:             *theMat,
		pileChooseCount: 0,
	}, nil
}

func pileToSerializablePile(pile deck.Pile) Pile {
	cards := make([]Card, 0)
	for _, card := range pile.Cards() {
		cards = append(cards, Card(card.String()))
	}
	return cards
}

func splitIntoThreePiles(sample deck.Pile) (*mat.Mat, error) {
	theMat := mat.New()
	for {
		var card deck.Card
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
