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
	var piles []deck.Pile
	for _, pile := range m.piles {
		piles = append(piles, pile)
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
