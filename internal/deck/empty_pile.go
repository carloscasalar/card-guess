package deck

import "github.com/carloscasalar/card-guess/pkg/threepilestrick"

type emptyPile struct{}

func (e emptyPile) DrawCard() (threepilestrick.Card, threepilestrick.Pile, error) {
	return nil, e, ErrNoMoreCardsInThePile
}

func (e emptyPile) AddCard(card threepilestrick.Card) threepilestrick.Pile {
	return &pile{
		topCard:    card,
		otherCards: emptyPile{},
	}
}

func (e emptyPile) StackOnTopOf(pile threepilestrick.Pile) threepilestrick.Pile {
	return pile
}

func (e emptyPile) Cards() []threepilestrick.Card {
	return []threepilestrick.Card{}
}

func (e emptyPile) Size() int {
	return 0
}

func (e emptyPile) String() string {
	return ""
}
