package trick

import "github.com/carloscasalar/card-guess/internal/deck"

type Mat struct {
	piles        map[holder]deck.Pile
	incomingPile holder
}

func (m *Mat) PlaceIntoNextPile(card deck.Card) {
	m.piles[m.incomingPile] = m.piles[m.incomingPile].AddCard(card)
	m.incomingPile = m.incomingPile.NextPile()
}

func (m *Mat) Piles() []deck.Pile {
	var piles = make([]deck.Pile, len(m.piles))
	for i, pile := range m.piles {
		piles[i] = pile
	}
	return piles
}

func NewMat() *Mat {
	return &Mat{
		piles: map[holder]deck.Pile{
			firstPile:  deck.NewPile(),
			secondPile: deck.NewPile(),
			thirdPile:  deck.NewPile(),
		},
		incomingPile: firstPile,
	}
}
