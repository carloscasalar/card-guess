package deck

type demoDealer struct {
	wrappedDealer Dealer
}

func NewDemoDealer() Dealer {
	return &demoDealer{
		wrappedDealer: NewDealer(),
	}
}

func (d *demoDealer) ShuffleCards() {
	// Pretending shuffle cards
}

func (d *demoDealer) Deal() (Card, error) {
	return d.wrappedDealer.Deal()
}
