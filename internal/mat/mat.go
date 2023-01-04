package mat

import (
	"github.com/carloscasalar/card-guess/internal/deck"
	"github.com/carloscasalar/card-guess/pkg/threepilestrick"
)

type Mat interface {
	PlaceIntoNextPile(card threepilestrick.Card) Mat
	JoinWithPileInTheMiddle(holder PileHolder) threepilestrick.Pile
	FirstPile() threepilestrick.Pile
	SecondPile() threepilestrick.Pile
	ThirdPile() threepilestrick.Pile
}

func New() Mat {
	return regularMat{
		piles: map[PileHolder]threepilestrick.Pile{
			FirstPile:  deck.NewPile(),
			SecondPile: deck.NewPile(),
			ThirdPile:  deck.NewPile(),
		},
	}
}

type regularMat struct {
	piles map[PileHolder]threepilestrick.Pile
}

func (m regularMat) PlaceIntoNextPile(card threepilestrick.Card) Mat {
	theMat := m.copy()
	nextPile := theMat.nextPile()
	theMat.piles[nextPile] = theMat.piles[nextPile].AddCard(card)
	return theMat
}

func (m regularMat) JoinWithPileInTheMiddle(holder PileHolder) threepilestrick.Pile {
	firstHolder := holder.nextPile()
	pile := m.piles[firstHolder].StackOnTopOf(m.piles[holder])
	lastHolder := firstHolder.nextPile()
	pile = pile.StackOnTopOf(m.piles[lastHolder])
	return pile
}

func (m regularMat) FirstPile() threepilestrick.Pile {
	return m.piles[FirstPile]
}

func (m regularMat) SecondPile() threepilestrick.Pile {
	return m.piles[SecondPile]
}

func (m regularMat) ThirdPile() threepilestrick.Pile {
	return m.piles[ThirdPile]
}

func (m regularMat) copy() regularMat {
	piles := make(map[PileHolder]threepilestrick.Pile)
	piles[FirstPile] = m.piles[FirstPile]
	piles[SecondPile] = m.piles[SecondPile]
	piles[ThirdPile] = m.piles[ThirdPile]

	return regularMat{piles}
}

func (m regularMat) nextPile() PileHolder {
	lessCardPile := ThirdPile
	minPileSize := m.piles[ThirdPile].Size()

	holder := FirstPile
	for {
		pileSize := m.piles[holder].Size()
		if pileSize <= minPileSize {
			lessCardPile = holder
			break
		}
		holder = holder.nextPile()
	}
	return lessCardPile
}
