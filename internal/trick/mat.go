package trick

import "github.com/carloscasalar/card-guess/internal/deck"

type Mat struct {
	piles        map[PileHolder]deck.Pile
	incomingPile PileHolder
}

func (m Mat) PlaceIntoNextPile(card deck.Card) Mat {
	mat := m.copy()
	mat.piles[m.incomingPile] = mat.piles[m.incomingPile].AddCard(card)
	mat.incomingPile = m.incomingPile.nextPile()
	return mat
}

func (m Mat) JoinWithPileInTheMiddle(holder PileHolder) deck.Pile {
	firstHolder := holder.nextPile()
	pile := m.piles[firstHolder].StackOnTopOf(m.piles[holder])
	lastHolder := firstHolder.nextPile()
	pile = pile.StackOnTopOf(m.piles[lastHolder])
	return pile
}

func (m Mat) Piles() []PileInMat {
	var piles = make([]PileInMat, len(m.piles))
	for holder, pile := range m.piles {
		piles[holder] = PileInMat{holder, pile}
	}
	return piles
}

func (m Mat) GetPile(holder PileHolder) deck.Pile {
	return m.piles[holder]
}

func (m Mat) copy() Mat {
	piles := make(map[PileHolder]deck.Pile)
	piles[FirstPile] = m.piles[FirstPile]
	piles[SecondPile] = m.piles[SecondPile]
	piles[ThirdPile] = m.piles[ThirdPile]

	incomingPile := m.incomingPile
	return Mat{piles, incomingPile}
}

func NewMat() Mat {
	return Mat{
		piles: map[PileHolder]deck.Pile{
			FirstPile:  deck.NewPile(),
			SecondPile: deck.NewPile(),
			ThirdPile:  deck.NewPile(),
		},
		incomingPile: FirstPile,
	}
}
