package deck

type emptyPile struct{}

func (e emptyPile) DrawCard() (Card, Pile, error) {
	return nil, e, ErrNoMoreCardsInThePile
}

func (e emptyPile) AddCard(card Card) Pile {
	return &pile{
		topCard:    card,
		otherCards: emptyPile{},
	}
}

func (e emptyPile) StackOnTopOf(pile Pile) Pile {
	return pile
}

func (e emptyPile) Cards() []Card {
	return []Card{}
}

func (e emptyPile) Size() int {
	return 0
}

func (e emptyPile) String() string {
	return ""
}
