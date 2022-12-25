package deck

type Pile struct {
	card *Card
}

func (p *Pile) DrawCard() (*Card, error) {
	if p.card != nil {
		card := p.card
		p.card = nil
		return card, nil
	}
	return nil, NoMoreCardsInThePile
}

func NewPile(cards ...Card) *Pile {
	if len(cards) == 0 {
		return &Pile{}
	}
	return &Pile{&cards[0]}
}
