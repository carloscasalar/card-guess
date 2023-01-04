package deck

import (
	"math/rand"
	"time"

	"github.com/carloscasalar/card-guess/pkg/threepilestrick"
)

type Dealer interface {
	ShuffleCards()
	Deal() (threepilestrick.Card, error)
}

func NewDealer() Dealer {
	return &dealer{
		deck: newDeck(),
	}
}

type dealer struct {
	deck Pile
}

func (d *dealer) ShuffleCards() {
	rand.Seed(time.Now().UnixNano())
	cards := d.deck.Cards()
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	d.deck = NewPile(cards...)
}

func (d *dealer) Deal() (threepilestrick.Card, error) {
	card, resultingDeck, err := d.deck.DrawCard()
	if err != nil {
		return nil, err
	}
	d.deck = resultingDeck
	return card, nil
}
