package threepilestrick

import (
	"fmt"

	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/mat"
)

const trickSampleSize = 21

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
	return trickState{
		sample: sample,
		mat:    *theMat,
	}, nil
}

type trickState struct {
	sample deck.Pile
	mat    mat.Mat
}

func (i trickState) GuessMyCard() (*Card, error) {
	return nil, StillCannotGuessYourCard
}

func (i trickState) Sample() Pile {
	return pileToSerializablePile(i.sample)
}

func (i trickState) FirstPile() Pile {
	return pileToSerializablePile(i.mat.FirstPile())
}

func (i trickState) SecondPile() Pile {
	return pileToSerializablePile(i.mat.SecondPile())
}

func (i trickState) ThirdPile() Pile {
	return pileToSerializablePile(i.mat.ThirdPile())
}

func (i trickState) MyCardIsInPile(holder PileHolder) (Trick, error) {
	allCardsWithChosenPileInTheMiddle := i.mat.JoinWithPileInTheMiddle(mat.PileHolder(holder))
	newMat, err := splitIntoThreePiles(allCardsWithChosenPileInTheMiddle)
	if err != nil {
		return nil, err
	}
	return trickState{
		sample: i.sample,
		mat:    *newMat,
	}, nil
}

func pileToSerializablePile(pile deck.Pile) Pile {
	var cards Pile
	for _, card := range pile.Cards() {
		cards = append(cards, Card(card.String()))
	}
	return cards
}
