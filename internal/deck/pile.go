package deck

type Pile struct {
	firstCard  *Card
	otherCards *Pile
}

func (p *Pile) DrawCard() (*Card, error) {
	if p.firstCard != nil {
		card := p.firstCard
		p.firstCard, _ = p.otherCards.DrawCard()
		return card, nil
	}
	return nil, NoMoreCardsInThePile
}

func NewPile(cards ...Card) *Pile {
	if len(cards) == 0 {
		return &Pile{}
	}
	firstCard := &cards[0]
	otherCards := NewPile(cards[1:]...)
	return &Pile{firstCard, otherCards}
}
