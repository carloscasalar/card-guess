package deck

import dealerlib "github.com/carloscasalar/go-cards/v2/pkg/dealer"

type Dealer interface {
	ShuffleCards()
	Deal() (Card, error)
}

func NewDealer() Dealer {
	return &dealer{
		wrappedDealer: dealerlib.NewDealer(1),
	}
}

type dealer struct {
	wrappedDealer *dealerlib.Dealer
}

func (d *dealer) ShuffleCards() {
	d.wrappedDealer.ShuffleCards()
}

func (d *dealer) Deal() (Card, error) {
	card, err := d.wrappedDealer.Deal()
	if err != nil {
		return nil, err
	}

	var cardDealt Card = *card
	return cardDealt, nil
}
