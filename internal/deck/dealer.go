package deck

import "github.com/carloscasalar/go-cards/v2/pkg/dealer"

type Dealer interface {
	ShuffleCards()
	Deal() (Card, error)
}

func NewDealer(numberOfDecks uint8) Dealer {
	return &dealerAdapter{
		wrappedDealer: dealer.NewDealer(numberOfDecks),
	}
}

type dealerAdapter struct {
	wrappedDealer *dealer.Dealer
}

func (d *dealerAdapter) ShuffleCards() {
	d.wrappedDealer.ShuffleCards()
}

func (d *dealerAdapter) Deal() (Card, error) {
	card, err := d.wrappedDealer.Deal()
	if err != nil {
		return nil, err
	}

	var cardDealt Card = *card
	return cardDealt, nil
}
