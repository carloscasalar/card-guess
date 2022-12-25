package deck

type Pile struct {
	firstCard  *Card
	otherCards *Pile
}

func NewPile(cards ...Card) *Pile {
	if len(cards) == 0 {
		return &Pile{}
	}
	firstCard := &cards[0]
	otherCards := NewPile(cards[1:]...)
	return &Pile{firstCard, otherCards}
}

func (p *Pile) DrawCard() (*Card, error) {
	if p.firstCard != nil {
		card := p.firstCard
		p.firstCard, _ = p.otherCards.DrawCard()
		return card, nil
	}
	return nil, ErrNoMoreCardsInThePile
}

func (p *Pile) AddCard(card Card) {
	if p.firstCard == nil {
		p.firstCard = &card
		return
	}
	if p.otherCards == nil {
		p.otherCards = NewPile(*p.firstCard)
		p.firstCard = &card
		return
	}
	p.otherCards.AddCard(*p.firstCard)
	p.firstCard = &card
}

func (p *Pile) Cards() []Card {
	if p.firstCard == nil {
		return []Card{}
	}
	if p.otherCards == nil {
		return []Card{*p.firstCard}
	}
	return append([]Card{*p.firstCard}, p.otherCards.Cards()...)
}
