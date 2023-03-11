package threepilestrick

import (
	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/mat"
)

type finalTrickState struct {
	sample                deck.Pile
	mat                   mat.Mat
	pileWhereChosenCardIs mat.PileHolder
}

func (f finalTrickState) Sample() Pile {
	return pileToSerializablePile(f.sample)
}

func (f finalTrickState) FirstPile() Pile {
	return pileToSerializablePile(f.mat.FirstPile())
}

func (f finalTrickState) SecondPile() Pile {
	return pileToSerializablePile(f.mat.SecondPile())
}

func (f finalTrickState) ThirdPile() Pile {
	return pileToSerializablePile(f.mat.ThirdPile())
}

func (f finalTrickState) GuessMyCard() (*Card, error) {
	const positionWhereCardIsAfterThreeSplits = 4
	card, err := f.mat.CardFromPile(f.pileWhereChosenCardIs, positionWhereCardIsAfterThreeSplits)
	if err != nil {
		return nil, err
	}
	theCard := Card(card.String())
	return &theCard, nil
}

func (f finalTrickState) MyCardIsInPile(holder PileHolder) (Trick, error) {
	if mat.PileHolder(holder) != f.pileWhereChosenCardIs {
		return nil, ErrAskMeToGuessTheCardInstead
	}
	return f, nil
}
