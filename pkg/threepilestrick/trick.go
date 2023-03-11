package threepilestrick

import (
	"fmt"

	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/internal/mat"
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
	return trickState{
		sample:          sample,
		mat:             *theMat,
		pileChooseCount: 0,
	}, nil
}

type trickState struct {
	sample          deck.Pile
	mat             mat.Mat
	pileChooseCount int
}

func (i trickState) GuessMyCard() (*Card, error) {
	return nil, ErrStillCannotGuessYourCard
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
	return trickState{
		sample:          i.sample,
		mat:             *newMat,
		pileChooseCount: pileChooseCount,
	}, nil
}

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

func pileToSerializablePile(pile deck.Pile) Pile {
	cards := make([]Card, 0)
	for _, card := range pile.Cards() {
		cards = append(cards, Card(card.String()))
	}
	return cards
}
