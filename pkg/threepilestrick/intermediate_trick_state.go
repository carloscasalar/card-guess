package threepilestrick

import (
	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/mat"
)

type intermediateTrickState struct {
	sample          deck.Pile
	mat             mat.Mat
	pileChooseCount int
}

func (i intermediateTrickState) GuessMyCard() (*Card, error) {
	return nil, ErrStillCannotGuessYourCard
}

func (i intermediateTrickState) Sample() Pile {
	return pileToSerializablePile(i.sample)
}

func (i intermediateTrickState) FirstPile() Pile {
	return pileToSerializablePile(i.mat.FirstPile())
}

func (i intermediateTrickState) SecondPile() Pile {
	return pileToSerializablePile(i.mat.SecondPile())
}

func (i intermediateTrickState) ThirdPile() Pile {
	return pileToSerializablePile(i.mat.ThirdPile())
}

func (i intermediateTrickState) MyCardIsInPile(holder PileHolder) (Trick, error) {
	const requiredChooseRounds = 3
	pileChooseCount := i.pileChooseCount + 1
	chosenPileHolder := mat.PileHolder(holder)
	if pileChooseCount == requiredChooseRounds {
		return &finalTrickState{sample: i.sample, mat: i.mat, pileWhereChosenCardIs: chosenPileHolder}, nil
	}
	allCardsWithChosenPileInTheMiddle := i.mat.JoinWithPileInTheMiddle(chosenPileHolder)
	newMat, err := splitIntoThreePiles(allCardsWithChosenPileInTheMiddle)
	if err != nil {
		return nil, err
	}
	return intermediateTrickState{
		sample:          i.sample,
		mat:             *newMat,
		pileChooseCount: pileChooseCount,
	}, nil
}
