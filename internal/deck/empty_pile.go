package deck

type emptyPile struct{}

func (e emptyPile) DrawCard() (Card, Pile, error) {
	return nil, e, ErrNoMoreCardsInThePile
}

func (e emptyPile) AddCard(card Card) Pile {
	return &pile{
		firstCard:  card,
		otherCards: emptyPile{},
	}
}

func (e emptyPile) Cards() []Card {
	return []Card{}
}

func (e emptyPile) String() string {
	return ""
}
