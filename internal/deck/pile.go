package deck

type Pile struct{}

func (p *Pile) DrawCard() (*Card, error) {
	return nil, NoMoreCardsInThePile
}

func NewPile() *Pile {
	return &Pile{}
}
