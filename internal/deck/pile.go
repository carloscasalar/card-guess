package deck

import (
	"errors"
	"fmt"
	"strings"

	"github.com/carloscasalar/card-guess/pkg/threepilestrick"
)

func NewPile(cards ...threepilestrick.Card) threepilestrick.Pile {
	if len(cards) == 0 {
		return &emptyPile{}
	}
	firstCard := cards[0]
	otherCards := NewPile(cards[1:]...)
	return &pile{firstCard, otherCards}
}

type pile struct {
	topCard    threepilestrick.Card
	otherCards threepilestrick.Pile
}

func (p pile) DrawCard() (threepilestrick.Card, threepilestrick.Pile, error) {
	newFirstCard, newOtherCards, err := p.otherCards.DrawCard()
	if err != nil {
		if errors.Is(err, ErrNoMoreCardsInThePile) {
			return p.topCard, emptyPile{}, nil
		}
		return nil, p, err
	}
	drawnCard := p.topCard
	resultingPile := &pile{newFirstCard, newOtherCards}
	return drawnCard, resultingPile, nil
}

func (p pile) AddCard(card threepilestrick.Card) threepilestrick.Pile {
	return &pile{
		topCard:    card,
		otherCards: p,
	}
}

func (p pile) StackOnTopOf(otherPile threepilestrick.Pile) threepilestrick.Pile {
	cards := append(p.Cards(), otherPile.Cards()...)

	return NewPile(cards...)
}

func (p pile) Cards() []threepilestrick.Card {
	return append([]threepilestrick.Card{p.topCard}, p.otherCards.Cards()...)
}

func (p pile) Size() int {
	return p.otherCards.Size() + 1
}

func (p pile) String() string {
	var cardStrings = make([]string, p.Size())
	for i, card := range p.Cards() {
		cardStrings[i] = fmt.Sprintf("%v", card)
	}
	return strings.Join(cardStrings, "  ")
}
