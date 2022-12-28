package trick

import "github.com/carloscasalar/card-guess/internal/deck"

type Mat struct {
	piles        map[PileHolder]deck.Pile
	incomingPile PileHolder
}

func (m Mat) PlaceIntoNextPile(card deck.Card) Mat {
	newPiles := make(map[PileHolder]deck.Pile)

	currentPile := m.incomingPile
	nextPile := currentPile.nextPile()
	lastPile := nextPile.nextPile()

	newPiles[currentPile] = m.piles[currentPile].AddCard(card)
	newPiles[nextPile] = m.piles[nextPile]
	newPiles[lastPile] = m.piles[lastPile]

	return Mat{
		piles:        newPiles,
		incomingPile: nextPile,
	}
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
