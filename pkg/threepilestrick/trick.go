package threepilestrick

import (
	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/trick"
)

type Card string
type Pile []Card

type Trick interface {
	Sample() Pile
	FirstPile() Pile
	SecondPile() Pile
	ThirdPile() Pile
	NextStep() CommandType
}

func New(shuffleBeforeInitialDraw bool) (Trick, error) {
	theTrick, err := trick.New(shuffleBeforeInitialDraw)
	if err != nil {
		return nil, err
	}
	return initialState{theTrick}, nil
}

type initialState struct {
	trick trick.Trick
}

func (i initialState) NextStep() CommandType {
	return ChoosePileWhereYourCardIs
}

func (i initialState) Sample() Pile {
	return pileToSerializablePile(i.trick.Sample())
}

func (i initialState) FirstPile() Pile {
	return pileToSerializablePile(i.trick.Mat().FirstPile())
}

func (i initialState) SecondPile() Pile {
	return pileToSerializablePile(i.trick.Mat().SecondPile())
}

func (i initialState) ThirdPile() Pile {
	return pileToSerializablePile(i.trick.Mat().ThirdPile())
}

func pileToSerializablePile(pile deck.Pile) Pile {
	var cards Pile
	for _, card := range pile.Cards() {
		cards = append(cards, Card(card.String()))
	}
	return cards
}
