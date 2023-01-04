package mat

import "github.com/carloscasalar/card-guess/internal/deck"

type Mat interface {
	PlaceIntoNextPile(card deck.Card) Mat
	JoinWithPileInTheMiddle(holder PileHolder) deck.Pile
	Piles() []PileInMat
	Pile(holder PileHolder) deck.Pile
}

func New() Mat {
	return regularMat{
		piles: map[PileHolder]deck.Pile{
			FirstPile:  deck.NewPile(),
			SecondPile: deck.NewPile(),
			ThirdPile:  deck.NewPile(),
		},
	}
}

type regularMat struct {
	piles map[PileHolder]deck.Pile
}

func (m regularMat) PlaceIntoNextPile(card deck.Card) Mat {
	theMat := m.copy()
	nextPile := theMat.nextPile()
	theMat.piles[nextPile] = theMat.piles[nextPile].AddCard(card)
	return theMat
}

func (m regularMat) JoinWithPileInTheMiddle(holder PileHolder) deck.Pile {
	firstHolder := holder.nextPile()
	pile := m.piles[firstHolder].StackOnTopOf(m.piles[holder])
	lastHolder := firstHolder.nextPile()
	pile = pile.StackOnTopOf(m.piles[lastHolder])
	return pile
}

func (m regularMat) Piles() []PileInMat {
	var piles = make([]PileInMat, len(m.piles))
	for holder, pile := range m.piles {
		piles[holder] = PileInMat{holder, pile}
	}
	return piles
}

func (m regularMat) Pile(holder PileHolder) deck.Pile {
	return m.piles[holder]
}

func (m regularMat) copy() regularMat {
	piles := make(map[PileHolder]deck.Pile)
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
